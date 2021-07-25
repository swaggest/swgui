package swgui

// Config is used for Swagger UI handler configuration.
type Config struct {
	Title       string `json:"title"`          // title of index file
	SwaggerJSON string `json:"swaggerJsonUrl"` // url to swagger.json document specification
	BasePath    string `json:"basePath"`       // base url to docs

	ShowTopBar         bool              `json:"showTopBar"`         // Show navigation top bar, hidden by default
	JsonEditor         bool              `json:"jsonEditor"`         // Enable visual json editor support (experimental, can fail with complex schemas)
	PreAuthorizeApiKey map[string]string `json:"preAuthorizeApiKey"` // Map of security name to key value

	// SettingsUI contains keys and plain javascript values of SwaggerUIBundle configuration.
	// Overrides default values.
	// See https://swagger.io/docs/open-source-tools/swagger-ui/usage/configuration/ for available options.
	SettingsUI map[string]string `json:"-"`
}
