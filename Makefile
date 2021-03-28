SWAGGER_UI_VERSION := v3.45.1

update:
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION)/dist/swagger-ui-bundle.js -o ./v3/static/swagger-ui-bundle.js
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION)/dist/swagger-ui-standalone-preset.js -o ./v3/static/swagger-ui-standalone-preset.js
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION)/dist/swagger-ui.js -o ./v3/static/swagger-ui.js
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION)/dist/swagger-ui.css -o ./v3/static/swagger-ui.css
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION)/dist/oauth2-redirect.html -o ./v3/static/oauth2-redirect.html
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION)/dist/favicon-32x32.png -o ./v3/static/favicon-32x32.png
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/$(SWAGGER_UI_VERSION)/dist/favicon-16x16.png -o ./v3/static/favicon-16x16.png
	rm -rf ./v3/static/*.gz
	go run ./v3/gen/gen.go
	zopfli --i50 ./v3/static/*.js && rm -f ./v3/static/*.js
	zopfli --i50 ./v3/static/*.css && rm -f ./v3/static/*.css
	zopfli --i50 ./v3/static/*.html && rm -f ./v3/static/*.html
