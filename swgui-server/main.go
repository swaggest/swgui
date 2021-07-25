package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	v3 "github.com/swaggest/swgui/v3"
)

var (
	host  string
	port  uint
	title string
)

func main() {
	flag.StringVar(&host, "host", "", "Host")
	flag.UintVar(&port, "port", 8080, "Port")
	flag.StringVar(&title, "title", "API Document", "Page title")

	flag.Parse()

	var h http.Handler = v3.NewHandler(title, "/swagger.json", "/")

	http.Handle("/", h)
	fmt.Printf("Listening at %s:%d\n", host, port)

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
