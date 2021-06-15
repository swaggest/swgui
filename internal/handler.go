package internal

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strings"

	"github.com/swaggest/swgui"
)

// Handler handles swagger UI request.
type Handler struct {
	swgui.Config

	ConfigJson template.JS

	tpl          *template.Template
	staticServer http.Handler
}

// NewHandlerWithConfig returns a HTTP handler for swagger UI.
func NewHandlerWithConfig(config swgui.Config, assetsBase, faviconBase string, staticServer http.Handler) *Handler {
	config.BasePath = strings.TrimSuffix(config.BasePath, "/") + "/"

	h := &Handler{
		Config: config,
	}
	j, _ := json.Marshal(h.Config)
	h.ConfigJson = template.JS(j)
	h.tpl, _ = template.New("index").Parse(IndexTpl(assetsBase, faviconBase, config))
	if staticServer != nil {
		h.staticServer = http.StripPrefix(h.BasePath, staticServer)
	}
	return h
}

// ServeHTTP implements http.Handler interface to handle swagger UI request.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.TrimSuffix(r.URL.Path, "/") != strings.TrimSuffix(h.BasePath, "/") && h.staticServer != nil {
		h.staticServer.ServeHTTP(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err := h.tpl.Execute(w, h); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
