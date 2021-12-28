//go:build swguicdn
// +build swguicdn

package v4emb

import (
	"net/http"

	"github.com/swaggest/swgui/v4cdn"
)

var staticServer http.Handler

const (
	assetsBase  = v4cdn.AssetsBase
	faviconBase = v4cdn.FaviconBase
)
