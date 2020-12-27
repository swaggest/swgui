// +build go1.16,!swguicdn

package v3

import (
	"github.com/swaggest/swgui/statigz"
	"github.com/swaggest/swgui/statigz/brotli"
	"github.com/swaggest/swgui/v3/static"
)

//var staticServer = gzipped.FileServer(http.FS(static.FS))
var staticServer = statigz.FileServer(static.FS, brotli.AddEncoding())

const (
	assetsBase  = "{{ .BasePath }}"
	faviconBase = "{{ .BasePath }}"
)
