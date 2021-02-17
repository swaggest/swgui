package v3emb

import (
	"github.com/swaggest/swgui"
	"github.com/swaggest/swgui/internal"
)

// Handler handles swagger UI request.
type Handler = internal.Handler

// NewHandler returns a HTTP handler for swagger UI.
func NewHandler(title, swaggerJSONPath string, basePath string) *Handler {
	return NewHandlerWithConfig(swgui.Config{
		Title:       title,
		SwaggerJSON: swaggerJSONPath,
		BasePath:    basePath,
	})
}

// NewHandlerWithConfig returns a HTTP handler for swagger UI.
func NewHandlerWithConfig(config swgui.Config) *Handler {
	return internal.NewHandlerWithConfig(config, assetsBase, faviconBase, staticServer)
}
