//go:build swguicdn
// +build swguicdn

package v5emb

import (
	"net/http"
)

var staticServer http.Handler

const (
	assetsBase  = v5cdn.AssetsBase
	faviconBase = v5cdn.FaviconBase
)
