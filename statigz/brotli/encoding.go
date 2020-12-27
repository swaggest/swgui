package brotli

import (
	"github.com/andybalholm/brotli"
	"github.com/swaggest/swgui/statigz"
	"io"
)

// AddEncoding is an option that prepends brotli to encodings of statigz.Server.
//
// It is located in a separate package to allow better control of imports graph.
func AddEncoding() func(server *statigz.Server) {
	return func(server *statigz.Server) {
		enc := statigz.Encoding{
			FileExt:         ".br",
			ContentEncoding: "br",
			Decoder: func(r io.Reader) (io.Reader, error) {
				return brotli.NewReader(r), nil
			},
		}

		server.Encodings = append([]statigz.Encoding{enc}, server.Encodings...)
	}
}
