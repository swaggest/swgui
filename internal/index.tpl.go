package internal

import (
	"sort"
	"strings"

	"github.com/swaggest/swgui"
)

// IndexTpl creates page template.
//
//nolint:funlen // The template is long.
func IndexTpl(assetsBase, faviconBase string, cfg swgui.Config) string {
	settings := map[string]string{
		"url":         "url",
		"dom_id":      "'#swagger-ui'",
		"deepLinking": "true",
		"presets": `[
				SwaggerUIBundle.presets.apis,
				SwaggerUIStandalonePreset
			]`,
		"plugins": `[
				SwaggerUIBundle.plugins.DownloadUrl
			]`,
		"layout":                   `"StandaloneLayout"`,
		"showExtensions":           "true",
		"showCommonExtensions":     "true",
		"validatorUrl":             "null",
		"defaultModelsExpandDepth": "-1", // Hides schemas, override with value "1" in Config.SettingsUI to show schemas.
		`onComplete`: `function() {
                if (cfg.preAuthorizeApiKey) {
                    for (var name in cfg.preAuthorizeApiKey) {
                        ui.preauthorizeApiKey(name, cfg.preAuthorizeApiKey[name]);
                    }
                }

                var dom = document.querySelector('.scheme-container select');
                for (var key in dom) {
                    if (key.startsWith("__reactInternalInstance$")) {
                        var compInternals = dom[key]._currentElement;
                        var compWrapper = compInternals._owner;
                        compWrapper._instance.setScheme(window.location.protocol.slice(0,-1));
                    }
                }
            }`,
	}

	for k, v := range cfg.SettingsUI {
		settings[k] = v
	}

	settingsStr := make([]string, 0, len(settings))
	for k, v := range settings {
		settingsStr = append(settingsStr, "\t\t\t"+k+": "+v)
	}

	sort.Strings(settingsStr)

	return `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ .Title }} - Swagger UI</title>
    <link rel="stylesheet" type="text/css" href="` + assetsBase + `swagger-ui.css">
    <link rel="icon" type="image/png" href="` + faviconBase + `favicon-32x32.png" sizes="32x32"/>
    <link rel="icon" type="image/png" href="` + faviconBase + `favicon-16x16.png" sizes="16x16"/>
    <style>
        html {
            box-sizing: border-box;
            overflow: -moz-scrollbars-vertical;
            overflow-y: scroll;
        }

        *,
        *:before,
        *:after {
            box-sizing: inherit;
        }

        body {
            margin: 0;
            background: #fafafa;
        }
    </style>
</head>

<body>
<div id="swagger-ui"></div>

<script src="` + assetsBase + `swagger-ui-bundle.js"></script>
<script src="` + assetsBase + `swagger-ui-standalone-preset.js"></script>
<script>
    window.onload = function () {
        var cfg = {{ .ConfigJson }};
        var url = cfg.swaggerJsonUrl;
        if (!url.startsWith("https://") && !url.startsWith("http://")) {
            if (url.startsWith(".")) {
                var path = window.location.pathname;
                path = path.endsWith("/") ? path : path + "/";
                url = window.location.protocol + "//" + window.location.host + path + url;
            } else {
                url = window.location.protocol + "//" + window.location.host + url;
            }
        }

        // Build a system
        var settings = {
` + strings.Join(settingsStr, ",\n") + `
        };

        if (cfg.showTopBar == false) {
            settings.plugins.push(function () {
                return {
                    components: {
                        Topbar: function () {
                            return null;
                        }
                    }
                }
            });
        }

        if (cfg.hideCurl) {
            settings.plugins.push(() => {return {wrapComponents: {curl: () => () => null}}});
        }

        window.ui = SwaggerUIBundle(settings);
    }
</script>
</body>
</html>
`
}
