// Package v3 provides embedded Swagger UI v3 assets.
package v3

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

// NewWithConfig creates configurable handler constructor.
func NewWithConfig(config swgui.Config) func(title, swaggerJSONPath string, basePath string) http.Handler {
	return func(title, swaggerJSONPath string, basePath string) http.Handler {
		if config.Title == "" {
			config.Title = title
		}

		if config.SwaggerJSON == "" {
			config.SwaggerJSON = swaggerJSONPath
		}

		if config.BasePath == "" {
			config.BasePath = basePath
		}

		return NewHandlerWithConfig(config)
	}
}

// NewHandler creates HTTP handler for Swagger UI.
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
