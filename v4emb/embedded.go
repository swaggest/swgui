// +build go1.16,!swguicdn

package v4emb

import (
	"github.com/swaggest/swgui/v4/static"
	"github.com/vearutop/statigz"
)

var staticServer = statigz.FileServer(static.FS)

const (
	assetsBase  = "{{ .BasePath }}"
	faviconBase = "{{ .BasePath }}"
)
