package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/ryandotsmith/32k.io/net/http/limit"
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

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if req.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, req)
	})
}

func usetls(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xfp := r.Header.Get("x-forwarded-proto")
		if xfp == "http" {
			redir := "https://" + r.Host + r.RequestURI
			w.Header().Set("Connection", "close")
			http.Redirect(w, r, redir, http.StatusTemporaryRedirect)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	dir := flag.String("data", "./sdat", "data directory")
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

	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/c/suggest", suggestion)

	var handler http.Handler
	handler = limit.NewHandler(mux, 10)
	handler = limit.MaxBytes(handler, limit.OneMB)
	handler = cors(handler)
	handler = usetls(handler)

	var listen = ":8080"
	if l := os.Getenv("LISTEN"); l != "" {
		listen = l
	}
	fmt.Printf("listening on %q\n", listen)

	// Timeout settings based on Filippo's late-2016 blog post
	// https://blog.filippo.io/exposing-go-on-the-internet/.
	srv := &http.Server{
		Addr:        listen,
		ReadTimeout: 5 * time.Second,
		// must be higher than the event handler timeout (10s)
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      handler,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Fprintln(os.Stderr, "ListenAndServe:", err)
		os.Exit(1)
	}
}
