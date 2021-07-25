// +build swguicdn

package v3emb

import (
	"net/http"

	"github.com/swaggest/swgui/v3cdn"
)

var staticServer http.Handler

const (
	assetsBase  = v3cdn.AssetsBase
	faviconBase = v3cdn.FaviconBase
)
