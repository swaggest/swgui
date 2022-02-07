// Package v3emb provides Swagger UI v3 with Go embed.
package v3emb

import (
	"net/http"

	"github.com/swaggest/swgui"
	"github.com/swaggest/swgui/internal"
)

// Handler handles swagger UI request.
type Handler = internal.Handler

// New creates HTTP handler for Swagger UI.
func New(title, swaggerJSONPath string, basePath string) http.Handler {
	return NewHandler(title, swaggerJSONPath, basePath)
}

// NewHandler creates HTTP handler for Swagger UI.
func NewHandler(title, swaggerJSONPath string, basePath string) *Handler {
	return NewHandlerWithConfig(swgui.Config{
		Title:       title,
		SwaggerJSON: swaggerJSONPath,
		BasePath:    basePath,
	})
}

// NewHandlerWithConfig creates HTTP handler for Swagger UI.
func NewHandlerWithConfig(config swgui.Config) *Handler {
	return internal.NewHandlerWithConfig(config, assetsBase, faviconBase, staticServer)
}
