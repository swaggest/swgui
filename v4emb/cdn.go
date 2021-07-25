// +build swguicdn

package v4emb

import (
	"github.com/swaggest/swgui/v4cdn"
	"net/http"
)

var staticServer http.Handler

const (
	assetsBase  = v4cdn.AssetsBase
	faviconBase = v4cdn.FaviconBase
)
