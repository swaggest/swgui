// +build go1.16,!swguicdn

package v3

import (
	"compress/gzip"
	"embed"
	"hash/fnv"
	"io"
	"mime"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/andybalholm/brotli"
	"github.com/swaggest/swgui/v3/static"
)

//var staticServer = gzipped.FileServer(http.FS(static.FS))
var staticServer = staticCompressedServer(static.FS)

type fileInfo struct {
	hash string
	size int64
}

type compressedServer struct {
	onError func(rw http.ResponseWriter, r *http.Request, err error)
	info    map[string]fileInfo
	fs      embed.FS
}

// staticCompressedServer directly serves compressed files from embedded file system.
func staticCompressedServer(fs embed.FS) http.Handler {
	cs := compressedServer{
		fs:   fs,
		info: make(map[string]fileInfo),
		onError: func(rw http.ResponseWriter, r *http.Request, err error) {
			http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		},
	}

	// Reading from "." is not expected to fail.
	if err := cs.hashDir("."); err != nil {
		panic(err)
	}

	return &cs
}

func (cs *compressedServer) hashDir(p string) error {
	files, err := cs.fs.ReadDir(p)
	if err != nil {
		return err
	}

	for _, f := range files {
		fn := path.Join(p, f.Name())
		if f.IsDir() {
			if err = cs.hashDir(fn); err != nil {
				return err
			}
		}

		h := fnv.New64()

		f, err := cs.fs.Open(fn)
		if err != nil {
			return err
		}

		n, err := io.Copy(h, f)
		if err != nil {
			return err
		}

		cs.info[path.Clean(fn)] = fileInfo{
			hash: strconv.FormatUint(h.Sum64(), 36),
			size: n,
		}
	}

	return nil
}

func (cs *compressedServer) serve(rw http.ResponseWriter, req *http.Request, fn, suf, enc, hash string, cl int64,
	decompress func(r io.Reader) (io.Reader, error)) {
	if m := req.Header.Get("If-None-Match"); m == hash {
		rw.WriteHeader(http.StatusNotModified)
		return
	}

	if ctype := mime.TypeByExtension(filepath.Ext(fn)); ctype != "" {
		rw.Header().Set("Content-Type", ctype)
	}

	rw.Header().Set("Etag", hash)

	if enc != "" {
		rw.Header().Set("Content-Encoding", enc)
	}

	var r io.Reader

	r, err := cs.fs.Open(fn + suf)
	if err != nil {
		cs.onError(rw, req, err)
		return
	}

	if decompress != nil {
		r, err = decompress(r)
		if err != nil {
			cs.onError(rw, req, err)
			return
		}
	} else {
		rw.Header().Set("Content-Length", strconv.Itoa(int(cl)))
	}

	if rs, ok := r.(io.ReadSeeker); ok {
		http.ServeContent(rw, req, fn, time.Time{}, rs)
		return
	}

	_, err = io.Copy(rw, r)
	if err != nil {
		cs.onError(rw, req, err)
		return
	}

	return
}

func (cs *compressedServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		rw.Header().Set("Allow", http.MethodGet)
		http.Error(rw, "Method Not Allowed\n\nmethod should be GET", http.StatusMethodNotAllowed)
		return
	}

	// Check accepted encodings.
	gzipAccepted := false
	brotliAccepted := false

	ae := req.Header.Get("Accept-Encoding")
	if ae != "" {
		ae = strings.ToLower(ae)

		n := strings.Index(ae, "gzip")
		if n >= 0 {
			gzipAccepted = true
		}

		n = strings.Index(ae, "br")
		if n >= 0 {
			brotliAccepted = true
		}
	}

	fn := strings.TrimPrefix(req.URL.Path, "/")

	// Copy brotli compressed data into response.
	brotliInfo, brotliFound := cs.info[fn+".br"]
	if brotliAccepted && brotliFound {
		cs.serve(rw, req, fn, ".br", "br", brotliInfo.hash, brotliInfo.size, nil)

		return
	}

	// Copy gzip compressed data to response.
	gzipInfo, gzipFound := cs.info[fn+".gz"]
	if gzipAccepted && gzipFound {
		cs.serve(rw, req, fn, ".gz", "gzip", gzipInfo.hash, gzipInfo.size, nil)

		return
	}

	// Copy uncompressed data to response.
	uncompressedInfo, uncompressedFound := cs.info[fn]
	if uncompressedFound {
		cs.serve(rw, req, fn, "", "", uncompressedInfo.hash, uncompressedInfo.size, nil)

		return
	}

	// Decompress gzip data into response.
	if gzipFound {
		cs.serve(rw, req, fn, ".gz", "", gzipInfo.hash+"U", 0, func(r io.Reader) (io.Reader, error) {
			return gzip.NewReader(r)
		})

		return
	}

	// Decompress brotli data into response.
	if brotliFound {
		cs.serve(rw, req, fn, ".br", "", gzipInfo.hash+"U", 0, func(r io.Reader) (io.Reader, error) {
			return brotli.NewReader(r), nil
		})

		return
	}

	http.NotFound(rw, req)
}

const (
	assetsBase  = "{{ .BasePath }}"
	faviconBase = "{{ .BasePath }}"
)
