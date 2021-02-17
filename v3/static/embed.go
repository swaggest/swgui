// +build go1.16

// Package static contains files to embed.
package static

import (
	"embed"
)

//go:embed *.png *.gz
var FS embed.FS
