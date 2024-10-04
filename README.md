# Swagger UI

[![GoDevDoc](https://img.shields.io/badge/dev-doc-00ADD8?logo=go)](https://pkg.go.dev/github.com/swaggest/swgui)

Package `swgui` (Swagger UI) provides HTTP handler to serve Swagger UI. All assets are embedded in Go source code, so
just build and run.

### V5

Static assets for `v5` are built from Swagger
UI [v5.17.14](https://github.com/swagger-api/swagger-ui/releases/tag/v5.17.14).

[CDN-based](https://cdnjs.com/libraries/swagger-ui) `v5cdn` uses Swagger
UI [v5.17.14](https://github.com/swagger-api/swagger-ui/releases/tag/v5.17.14).

### V4

Static assets for `v4` are built from Swagger
UI [v4.19.1](https://github.com/swagger-api/swagger-ui/releases/tag/v4.19.1).

[CDN-based](https://cdnjs.com/libraries/swagger-ui) `v4cdn` uses Swagger
UI [v4.18.3](https://github.com/swagger-api/swagger-ui/releases/tag/v4.18.3).

### V3

Static assets for `v3` are built from Swagger
UI [v3.52.5](https://github.com/swagger-api/swagger-ui/releases/tag/v3.52.5).

[CDN-based](https://cdnjs.com/libraries/swagger-ui) `v3cdn` uses Swagger
UI [v3.52.4](https://github.com/swagger-api/swagger-ui/releases/tag/v3.52.4).

## How to use

```go
package main

import (
	"net/http"

	"github.com/swaggest/swgui/v5emb"
)

func main() {
	http.Handle("/api1/docs/", v5emb.New(
		"Petstore",
		"https://petstore3.swagger.io/api/v3/openapi.json",
		"/api1/docs/",
	))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello World!"))
	})

	println("docs at http://localhost:8080/api1/docs/")
	_ = http.ListenAndServe("localhost:8080", http.DefaultServeMux)
}
```


If you use `go1.16` or later, you can import natively embedded assets with `"github.com/swaggest/swgui/v5emb"`, it may
help to lower application memory usage.

## Use CDN for assets

In order to reduce binary size you can import `github.com/swaggest/swgui/v3cdn` to use CDN hosted assets.

Also you can use `swguicdn` build tag to enable CDN mode for `github.com/swaggest/swgui/v3` import.

Be aware that CDN mode may be considered inappropriate for security or networking reasons.

## Documentation viewer CLI tool

Install `swgui`.

```
go install github.com/swaggest/swgui/cmd/swgui@latest
```

Or download binary from [releases](https://github.com/swaggest/swgui/releases).

### Linux AMD64

```
wget https://github.com/swaggest/swgui/releases/latest/download/linux_amd64.tar.gz && tar xf linux_amd64.tar.gz && rm linux_amd64.tar.gz
./swgui -version
```

### Macos Intel

```
wget https://github.com/swaggest/swgui/releases/latest/download/darwin_amd64.tar.gz && tar xf darwin_amd64.tar.gz && rm darwin_amd64.tar.gz
codesign -s - ./swgui
./swgui -version
```

### Macos Apple Silicon (M1, etc...)

```
wget https://github.com/swaggest/swgui/releases/latest/download/darwin_arm64.tar.gz && tar xf darwin_arm64.tar.gz && rm darwin_arm64.tar.gz
codesign -s - ./swgui
./swgui -version
```

Open spec file.

```
swgui my-openapi.yaml
```
