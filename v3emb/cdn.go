// +build swguicdn

package v3emb

import (
	"github.com/swaggest/swgui/v3cdn"
	"net/http"
)

var staticServer http.Handler

const (
	assetsBase  = v3cdn.AssetsBase
	faviconBase = v3cdn.FaviconBase
)
