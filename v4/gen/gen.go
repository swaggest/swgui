// Package main generates legacy static assets.
package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	var fs http.FileSystem = http.Dir("./v4/static")

	err := vfsgen.Generate(fs, vfsgen.Options{
		BuildTags:   "!swguicdn",
		PackageName: "v4",
		Filename:    "v4/static.go",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
