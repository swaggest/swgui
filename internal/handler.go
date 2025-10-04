// Package internal provides internal handler implementation.
package internal

import (
	"encoding/json"
	"html/template"
	"io"
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

	if h.InternalBasePath == "" {
		h.InternalBasePath = h.BasePath
	}

	h.InternalBasePath = strings.TrimSuffix(h.InternalBasePath, "/") + "/"

	j, err := json.Marshal(h.Config)
	if err != nil {
		panic(err)
	}

	h.ConfigJson = template.JS(j) //nolint:gosec // Data is well formed.

	h.tpl, err = template.New("index").Parse(IndexTpl(assetsBase, faviconBase, config))
	if err != nil {
		panic(err)
	}

	if staticServer != nil {
		h.staticServer = http.StripPrefix(h.InternalBasePath, staticServer)
	}

	return h
}

// ServeHTTP implements http.Handler interface to handle swagger UI request.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.TrimSuffix(r.URL.Path, "/") != strings.TrimSuffix(h.InternalBasePath, "/") && h.staticServer != nil {
		h.staticServer.ServeHTTP(w, r)

		return
	}

	if u := r.URL.Query().Get("proxy"); u != "" && h.Proxy {
		h.proxyRequest(u, w, r)

		return
	}

	w.Header().Set("Content-Type", "text/html")

	if err := h.tpl.Execute(w, h); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) proxyRequest(u string, w http.ResponseWriter, r *http.Request) {
	b := r.Body
	defer b.Close()

	req, err := http.NewRequest(r.Method, u, b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	req.Header = r.Header

	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	defer resp.Body.Close()

	hd := w.Header()

	for k, vv := range resp.Header {
		for _, v := range vv {
			hd.Add(k, v)
		}
	}

	w.WriteHeader(resp.StatusCode)

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
