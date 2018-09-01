package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/ryandotsmith/32k.io/net/http/limit"
	"github.com/ryandotsmith/32k.io/net/mylisten"
	"github.com/ryandotsmith/32k.io/net/mytls"
)

var (
	//protects access to suggestions file
	sugMu   sync.Mutex
	sugFile *os.File
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hi")
}

func suggestion(w http.ResponseWriter, r *http.Request) {
	sugMu.Lock()
	if _, err := sugFile.WriteString(string(r.FormValue("suggestion")) + "\n"); err != nil {
		sugMu.Unlock()
		http.Error(w, "write suggestion", http.StatusInternalServerError)
		log.Fatal(err)
	}
	sugMu.Unlock()
	fmt.Fprintln(w, "thanks")
}

func main() {
	listen := flag.String("listen", "localhost:7000", "listen `address` (if no LISTEN_FDS)")
	dir := flag.String("data", "./sdat", "data directory")
	flag.Parse()

	err := os.MkdirAll(*dir, 0700)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	sugFile, err = os.OpenFile(*dir+"/suggestions.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	l, r, err := mylisten.SystemdOr(*listen)
	if err != nil {
		fmt.Fprintln(os.Stderr, "listen:", err)
		os.Exit(1)
	}

	if r != nil {
		go func() {
			rSrv := &http.Server{
				ReadTimeout:  5 * time.Second,
				WriteTimeout: 5 * time.Second,
				Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
					w.Header().Set("Connection", "close")
					url := "https://" + req.Host + req.URL.String()
					http.Redirect(w, req, url, http.StatusMovedPermanently)
				}),
			}
			err := rSrv.Serve(r)
			panic(err)
		}()
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/c/suggest", suggestion)

	var handler http.Handler
	handler = limit.NewHandler(mux, 10)
	handler = limit.MaxBytes(handler, limit.OneMB)

	// Timeout settings based on Filippo's late-2016 blog post
	// https://blog.filippo.io/exposing-go-on-the-internet/.
	srv := &http.Server{
		ReadTimeout: 5 * time.Second,
		// must be higher than the event handler timeout (10s)
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      handler,
	}

	cfg, err := mytls.LocalOrLets(*dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "tls:", err)
		os.Exit(1)
	}

	err = srv.Serve(tls.NewListener(l, cfg))
	if err != nil {
		fmt.Fprintln(os.Stderr, "ListenAndServe:", err)
		os.Exit(1)
	}
}
