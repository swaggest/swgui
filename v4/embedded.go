// +build !swguicdn

package v4

import "github.com/shurcooL/httpgzip"

var staticServer = httpgzip.FileServer(assets, httpgzip.FileServerOptions{})

const (
	assetsBase  = "{{ .BasePath }}"
	faviconBase = "{{ .BasePath }}"
)
