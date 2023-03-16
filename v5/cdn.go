//go:build swguicdn
// +build swguicdn

package v5

import (
	"net/http"

	"github.com/swaggest/swgui/v5cdn"
)

var staticServer http.Handler

const (
	assetsBase  = v5cdn.AssetsBase
	faviconBase = v5cdn.FaviconBase
)
