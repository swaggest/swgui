// Package main provides CLI tool to inspect OpenAPI schemas with Swagger UI.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os/exec"
	"path"
	"runtime"

	"github.com/bool64/dev/version"
	swgui "github.com/swaggest/swgui/v5emb"
)

func main() {
	ver := flag.Bool("version", false, "Show version and exit.")
	flag.Parse()

	if *ver {
		fmt.Printf("%s, Swagger UI %s\n", version.Info().Version, "v5.20.5")

		return
	}

	if flag.NArg() < 1 {
		fmt.Println("Usage: swgui <path-to-schema>")
		flag.PrintDefaults()

		return
	}

	filePathToSchema := flag.Arg(0)
	urlToSchema := "/" + path.Base(filePathToSchema)

	swh := swgui.NewHandler(filePathToSchema, urlToSchema, "/")
	hh := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.URL.Path == urlToSchema {
			http.ServeFile(rw, r, filePathToSchema)

			return
		}

		swh.ServeHTTP(rw, r)
	})

	srv := httptest.NewServer(hh)

	log.Println("Starting Swagger UI server at", srv.URL)
	log.Println("Press Ctrl+C to stop")

	if err := open(srv.URL); err != nil {
		log.Println("open browser:", err.Error())
	}

	<-make(chan struct{})
}

// open opens the specified URL in the default browser of the user.
func open(url string) error {
	var (
		cmd  string
		args []string
	)

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}

	args = append(args, url)

	return exec.Command(cmd, args...).Start() //nolint:gosec
}
