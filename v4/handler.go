// Package v4 provides Swagger UI v4 assets.
package v4

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
