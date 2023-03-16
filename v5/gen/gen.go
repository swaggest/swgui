// Package main generates legacy static assets.
package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	var fs http.FileSystem = http.Dir("./v5/static")

	err := vfsgen.Generate(fs, vfsgen.Options{
		BuildTags:   "!swguicdn",
		PackageName: "v5",
		Filename:    "v5/static.go",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
