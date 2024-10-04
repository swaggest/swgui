#GOLANGCI_LINT_VERSION := "v1.51.1" # Optional configuration to pinpoint golangci-lint version.

# The head of Makefile determines location of dev-go to include standard targets.
GO ?= go
export GO111MODULE = on

ifneq "$(GOFLAGS)" ""
  $(info GOFLAGS: ${GOFLAGS})
endif

ifneq "$(wildcard ./vendor )" ""
  $(info Using vendor)
  modVendor =  -mod=vendor
  ifeq (,$(findstring -mod,$(GOFLAGS)))
      export GOFLAGS := ${GOFLAGS} ${modVendor}
  endif
  ifneq "$(wildcard ./vendor/github.com/bool64/dev)" ""
  	DEVGO_PATH := ./vendor/github.com/bool64/dev
  endif
endif

ifeq ($(DEVGO_PATH),)
	DEVGO_PATH := $(shell GO111MODULE=on $(GO) list ${modVendor} -f '{{.Dir}}' -m github.com/bool64/dev)
	ifeq ($(DEVGO_PATH),)
    	$(info Module github.com/bool64/dev not found, downloading.)
    	DEVGO_PATH := $(shell export GO111MODULE=on && $(GO) get github.com/bool64/dev && $(GO) list -f '{{.Dir}}' -m github.com/bool64/dev)
	endif
endif

export CGO_ENABLED = 0
BUILD_LDFLAGS=-s -w
BUILD_PKG = ./cmd/swgui
export RELEASE_TARGETS="darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64"

-include $(DEVGO_PATH)/makefiles/main.mk
-include $(DEVGO_PATH)/makefiles/lint.mk
-include $(DEVGO_PATH)/makefiles/reset-ci.mk
-include $(DEVGO_PATH)/makefiles/release-assets.mk
-include $(DEVGO_PATH)/makefiles/build.mk

# Add your custom targets here.

SWAGGER_UI_VERSION_V3 := v3.52.5
SWAGGER_UI_VERSION_V4 := v4.19.1
SWAGGER_UI_VERSION_V5 := v5.17.14

## Update assets for Swagger UI v3
update-v3:
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION_V3)/dist/swagger-ui-bundle.js -o ./v3/static/swagger-ui-bundle.js
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION_V3)/dist/swagger-ui-standalone-preset.js -o ./v3/static/swagger-ui-standalone-preset.js
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION_V3)/dist/swagger-ui.js -o ./v3/static/swagger-ui.js
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION_V3)/dist/swagger-ui.css -o ./v3/static/swagger-ui.css
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION_V3)/dist/oauth2-redirect.html -o ./v3/static/oauth2-redirect.html
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION_V3)/dist/favicon-32x32.png -o ./v3/static/favicon-32x32.png
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION_V3)/dist/favicon-16x16.png -o ./v3/static/favicon-16x16.png
	rm -rf ./v3/static/*.gz
	go run ./v3/gen/gen.go
	zopfli --i50 ./v3/static/*.js && rm -f ./v3/static/*.js
	zopfli --i50 ./v3/static/*.css && rm -f ./v3/static/*.css
	zopfli --i50 ./v3/static/*.html && rm -f ./v3/static/*.html

## Update assets for Swagger UI v4
update-v4:
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION_V4)/dist/swagger-ui-bundle.js -o ./v4/static/swagger-ui-bundle.js
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION_V4)/dist/swagger-ui-standalone-preset.js -o ./v4/static/swagger-ui-standalone-preset.js
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION_V4)/dist/swagger-ui.js -o ./v4/static/swagger-ui.js
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION_V4)/dist/swagger-ui.css -o ./v4/static/swagger-ui.css
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION_V4)/dist/oauth2-redirect.html -o ./v4/static/oauth2-redirect.html
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION_V4)/dist/favicon-32x32.png -o ./v4/static/favicon-32x32.png
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION_V4)/dist/favicon-16x16.png -o ./v4/static/favicon-16x16.png
	rm -rf ./v4/static/*.gz
	go run ./v4/gen/gen.go
	zopfli --i50 ./v4/static/*.js && rm -f ./v4/static/*.js
	zopfli --i50 ./v4/static/*.css && rm -f ./v4/static/*.css
	zopfli --i50 ./v4/static/*.html && rm -f ./v4/static/*.html

## Update assets for Swagger UI v5
update-v5:
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION_V5)/dist/swagger-ui-bundle.js -o ./v5/static/swagger-ui-bundle.js
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION_V5)/dist/swagger-ui-standalone-preset.js -o ./v5/static/swagger-ui-standalone-preset.js
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION_V5)/dist/swagger-ui.js -o ./v5/static/swagger-ui.js
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION_V5)/dist/swagger-ui.css -o ./v5/static/swagger-ui.css
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION_V5)/dist/oauth2-redirect.html -o ./v5/static/oauth2-redirect.html
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION_V5)/dist/favicon-32x32.png -o ./v5/static/favicon-32x32.png
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION_V5)/dist/favicon-16x16.png -o ./v5/static/favicon-16x16.png
	rm -rf ./v5/static/*.gz
	go run ./v5/gen/gen.go
	zopfli --i50 ./v5/static/*.js && rm -f ./v5/static/*.js
	zopfli --i50 ./v5/static/*.css && rm -f ./v5/static/*.css
	zopfli --i50 ./v5/static/*.html && rm -f ./v5/static/*.html
