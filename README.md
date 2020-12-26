# Swagger UI

[![GoDevDoc](https://img.shields.io/badge/dev-doc-00ADD8?logo=go)](https://pkg.go.dev/github.com/swaggest/swgui)

Package `swgui` (Swagger UI) provides HTTP handler to serve Swagger UI.
All assets are embedded in Go source code, so just build and run.

Static assets for `v3` are built from Swagger UI [v3.38.0](https://github.com/swagger-api/swagger-ui/releases/tag/v3.38.0).

[CDN-based](https://cdnjs.com/libraries/swagger-ui) `v3cdn` uses Swagger UI [v3.38.0](https://github.com/swagger-api/swagger-ui/releases/tag/v3.38.0).

## How to use

```go
package main

import (
    "net/http"

    "github.com/swaggest/swgui/v3"
)

func main() {
    http.Handle("/", v3.NewHandler("My API", "/swagger.json", "/"))
    http.ListenAndServe(":8080", nil)
}
```

## Use CDN for assets

In order to reduce binary size you can import `github.com/swaggest/swgui/v3cdn` to use CDN hosted assets.

Also you can use `swguicdn` build tag to enable CDN mode for `github.com/swaggest/swgui/v3` import.

Be aware that CDN mode may be considered inappropriate for security or networking reasons.

## Run as standalone server

Install `swgui-server`

    go get github.com/swaggest/swgui/...

Start server

    swgui-server -port 8080
