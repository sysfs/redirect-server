// vim: set noet ts=4 sw=4 tw=79 cc=80 list listchars=tab\:\ \ ,lead\:·,trail\:·:

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	help        bool
	redirectURL string
)

func init() {
	flag.StringVar(&redirectURL, "redirect-url", "", "URL to redirect to (required)")
	flag.BoolVar(&help, "help", false, "Show help message")

	flag.Parse()

	if help {
		fmt.Println("Usage: go run main.go [--redirect-url=<url>] ")
		fmt.Println("  --redirect-url=<url>: URL to redirect all requests to.")
		fmt.Println("  -h, --help: Display this help message.")
		os.Exit(0)
	}

	if redirectURL == "" {
		fmt.Println("Error: --redirect-url flag is required")
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	http.HandleFunc("/", RedirectServer)
	http.ListenAndServe(":8080", nil)
}

func RedirectServer(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	http.Redirect(w, r, redirectURL, http.StatusMovedPermanently)
}

// helper to log a message to stdout in Common Log Format
func logRequest(r *http.Request) {
	now := time.Now()

	clfLog := fmt.Sprintf("%s - - [%s] \"%s %s %s\" %d %d \"%s\" \"%s\"\n",
		r.RemoteAddr, now.Format("02/Jan/2006:15:04:05 -0700"), r.Method,
		r.RequestURI, r.Proto, http.StatusMovedPermanently, 0, r.Referer(), r.UserAgent(),
	)

	log.Print(clfLog)
}
