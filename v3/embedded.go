// +build !swguicdn

package v3

import "github.com/shurcooL/httpgzip"

var staticServer = httpgzip.FileServer(assets, httpgzip.FileServerOptions{})

const (
	assetsBase  = "{{ .BasePath }}"
	faviconBase = "{{ .BasePath }}"
)
