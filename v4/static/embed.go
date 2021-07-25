// +build go1.16

// Package static contains files to embed.
package static

import (
	"embed"
)

//go:embed *.png *.gz
// FS holds embedded static assets.
var FS embed.FS
