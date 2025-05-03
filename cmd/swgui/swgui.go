// Package main provides CLI tool to inspect OpenAPI schemas with Swagger UI.
package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"
	"path"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/bool64/dev/version"
	"github.com/swaggest/swgui"
	"github.com/swaggest/swgui/v5emb"
)

func main() {
	var (
		listen      string
		skipBrowser bool
		ver         bool
		proxy       bool
	)

	flag.StringVar(&listen, "listen", "127.0.0.1:0", "listen address, port 0 picks a free random port")
	flag.BoolVar(&skipBrowser, "s", false, "skip browser opening")
	flag.BoolVar(&ver, "version", false, "Show version and exit.")
	flag.BoolVar(&proxy, "proxy", false, "Proxy requests.")

	flag.Parse()

	if ver {
		fmt.Printf("%s, Swagger UI %s\n", version.Info().Version, "v5.21.0")

		return
	}

	if flag.NArg() < 1 {
		fmt.Println("Usage: swgui <path-to-schema>")
		flag.PrintDefaults()

		return
	}

	filePathToSchema := flag.Arg(0)
	urlToSchema := "/" + path.Base(filePathToSchema)

	cfg := swgui.Config{
		Title:       filePathToSchema,
		SwaggerJSON: urlToSchema,
		BasePath:    "/",
		Proxy:       proxy,
	}

	swh := v5emb.NewHandlerWithConfig(cfg)
	hh := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.URL.Path == urlToSchema {
			http.ServeFile(rw, r, filePathToSchema)

			return
		}

		swh.ServeHTTP(rw, r)
	})

	listener, err := net.Listen("tcp", listen)
	if err != nil {
		log.Fatalf("failed to start server: %s", err)

		return
	}

	addr := listener.Addr().String()

	srv := &http.Server{Handler: hh, ReadHeaderTimeout: time.Second}

	go func() {
		if err := srv.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Println("failed to listen ans serve: ", err.Error())
		}
	}()

	if strings.HasPrefix(listen, ":") {
		m, err := interfaces(false)
		if err != nil {
			log.Println("find network interfaces:", err)
		} else {
			for _, v := range m {
				addr = v + listen
			}
		}
	}

	log.Println("Starting Swagger UI server at http://" + addr)
	log.Println("Press Ctrl+C to stop")

	if !skipBrowser {
		if err := openBrowser("http://" + addr); err != nil && !strings.Contains(err.Error(), "executable file not found") {
			log.Println("failed to open browser", err.Error())
		}
	}

	<-make(chan struct{})
}

// openBrowser opens the specified URL in the default browser of the user.
func openBrowser(url string) error {
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

// interfaces returns a `name:ip` map of the suitable interfaces found.
func interfaces(listAll bool) ([]string, error) {
	names := make([]string, 0)

	ifaces, err := net.Interfaces()
	if err != nil {
		return names, err
	}

	re := regexp.MustCompile(`^(veth|br\-|docker|lo|EHC|XHC|bridge|gif|stf|p2p|awdl|utun|tun|tap)`)

	for _, iface := range ifaces {
		if !listAll && re.MatchString(iface.Name) {
			continue
		}

		if iface.Flags&net.FlagUp == 0 {
			continue
		}

		ip, err := findIP(iface)
		if err != nil {
			continue
		}

		names = append(names, ip)
	}

	return names, nil
}

// FindIP returns the IP address of the passed interface, and an error.
func findIP(iface net.Interface) (string, error) {
	var ip string

	addrs, err := iface.Addrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok {
			if ipnet.IP.IsLinkLocalUnicast() {
				continue
			}

			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()

				continue
			}
			// Use IPv6 only if an IPv4 hasn't been found yet.
			// This is eventually overwritten with an IPv4, if found (see above)
			if ip == "" {
				ip = "[" + ipnet.IP.String() + "]"
			}
		}
	}

	if ip == "" {
		return "", errors.New("unable to find an IP for this interface")
	}

	return ip, nil
}
