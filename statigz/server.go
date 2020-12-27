package statigz

import (
	"compress/gzip"
	"hash/fnv"
	"io"
	"io/fs"
	"mime"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Server is a http.Handler that directly serves compressed files from file system to capable agents.
//
// Please use FileServer to create an instance of Server.
//
// If agent does not accept encoding and uncompressed file is not available in file system,
// it would decompress the file before serving.
//
// Compressed files should have an additional extension to indicate their encoding,
// for example "style.css.gz" or "bundle.js.br".
//
// Behavior is similar to http://nginx.org/en/docs/http/ngx_http_gzip_static_module.html and
// https://github.com/lpar/gzipped, except compressed data can be decompressed for an incapable agent.
type Server struct {
	OnError   func(rw http.ResponseWriter, r *http.Request, err error)
	Encodings []Encoding

	info map[string]fileInfo
	fs   fs.ReadDirFS
}

// FileServer creates an instance of Server from file system.
//
// Typically file system would be an embed.FS.
//
//   //go:embed *.png *.br
//	 var FS embed.FS
func FileServer(fs fs.ReadDirFS, options ...func(server *Server)) *Server {
	s := Server{
		fs:   fs,
		info: make(map[string]fileInfo),
		OnError: func(rw http.ResponseWriter, r *http.Request, err error) {
			http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		},
		Encodings: []Encoding{GzipEncoding()},
	}

	for _, o := range options {
		o(&s)
	}

	// Reading from "." is not expected to fail.
	if err := s.hashDir("."); err != nil {
		panic(err)
	}

	return &s
}

func (s *Server) hashDir(p string) error {
	files, err := s.fs.ReadDir(p)
	if err != nil {
		return err
	}

	for _, f := range files {
		fn := path.Join(p, f.Name())
		if f.IsDir() {
			if err = s.hashDir(fn); err != nil {
				return err
			}
		}

		h := fnv.New64()

		f, err := s.fs.Open(fn)
		if err != nil {
			return err
		}

		n, err := io.Copy(h, f)
		if err != nil {
			return err
		}

		s.info[path.Clean(fn)] = fileInfo{
			hash: strconv.FormatUint(h.Sum64(), 36),
			size: n,
		}
	}

	return nil
}

func (s *Server) serve(rw http.ResponseWriter, req *http.Request, fn, suf, enc, hash string, cl int64,
	decompress func(r io.Reader) (io.Reader, error)) {
	if m := req.Header.Get("If-None-Match"); m == hash {
		rw.WriteHeader(http.StatusNotModified)

		return
	}

	ctype := mime.TypeByExtension(filepath.Ext(fn))
	if ctype == "" {
		ctype = "application/octet-stream" // Prevent unreliable Content-Type detection on compressed data.
	}

	rw.Header().Set("Content-Type", ctype)
	rw.Header().Set("Etag", hash)

	if enc != "" {
		rw.Header().Set("Content-Encoding", enc)
	}

	var r io.Reader

	r, err := s.fs.Open(fn + suf)
	if err != nil {
		s.OnError(rw, req, err)
		return
	}

	if cl > 0 {
		rw.Header().Set("Content-Length", strconv.Itoa(int(cl)))
	}

	if decompress != nil {
		r, err = decompress(r)
		if err != nil {
			s.OnError(rw, req, err)
			return
		}
	}

	if rs, ok := r.(io.ReadSeeker); ok {
		http.ServeContent(rw, req, fn, time.Time{}, rs)

		return
	}

	_, err = io.Copy(rw, r)
	if err != nil {
		s.OnError(rw, req, err)

		return
	}

	return
}

// ServeHTTP serves static files.
func (s *Server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		rw.Header().Set("Allow", http.MethodGet)
		http.Error(rw, "Method Not Allowed\n\nmethod should be GET", http.StatusMethodNotAllowed)
		return
	}

	fn := strings.TrimPrefix(req.URL.Path, "/")
	ae := req.Header.Get("Accept-Encoding")

	if ae != "" {
		ae = strings.ToLower(ae)

		for _, enc := range s.Encodings {
			if strings.Index(ae, enc.ContentEncoding) < 0 {
				continue
			}

			info, found := s.info[fn+enc.FileExt]
			if !found {
				continue
			}

			// Copy compressed data into response.
			s.serve(rw, req, fn, enc.FileExt, enc.ContentEncoding, info.hash, info.size, nil)

			return
		}
	}

	// Copy uncompressed data into response.
	uncompressedInfo, uncompressedFound := s.info[fn]
	if uncompressedFound {
		s.serve(rw, req, fn, "", "", uncompressedInfo.hash, uncompressedInfo.size, nil)

		return
	}

	// Decompress compressed data into response.
	for _, enc := range s.Encodings {
		info, found := s.info[fn+enc.FileExt]
		if !found || enc.Decoder == nil {
			continue
		}

		s.serve(rw, req, fn, enc.FileExt, "", info.hash+"U", 0, enc.Decoder)

		return
	}

	http.NotFound(rw, req)
}

// Encoding describes content encoding.
type Encoding struct {
	// FileExt is an extension of file with compressed content, for example ".gz".
	FileExt string

	// ContentEncoding is encoding name that is used in Accept-Encoding and Content-Encoding
	// headers, for example "gzip".
	ContentEncoding string

	// Decoder is a function that can decode data for an agent that does not accept encoding,
	// can be nil to disable dynamic decompression.
	Decoder func(r io.Reader) (io.Reader, error)
}

type fileInfo struct {
	hash string
	size int64
}

// OnError is an option to customize error handling in Server.
func OnError(onErr func(rw http.ResponseWriter, r *http.Request, err error)) func(server *Server) {
	return func(server *Server) {
		server.OnError = onErr
	}
}

// GzipEncoding provides gzip Encoding.
func GzipEncoding() Encoding {
	return Encoding{
		FileExt:         ".gz",
		ContentEncoding: "gzip",
		Decoder: func(r io.Reader) (io.Reader, error) {
			return gzip.NewReader(r)
		},
	}
}
