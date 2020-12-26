// +build swguicdn

package v3

import "net/http"

var staticServer http.Handler

const (
	assetsBase  = "https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.38.0/"
	faviconBase = "https://petstore.swagger.io/"
)
