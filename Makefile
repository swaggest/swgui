SWAGGER_UI_VERSION_V3 := v3.51.1
SWAGGER_UI_VERSION_V4 := v4.0.0-beta.1

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
