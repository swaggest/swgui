// Package main generates legacy assets.
package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	var fs http.FileSystem = http.Dir("./v3/static")

	err := vfsgen.Generate(fs, vfsgen.Options{
		BuildTags:   "!swguicdn",
		PackageName: "v3",
		Filename:    "v3/static.go",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
