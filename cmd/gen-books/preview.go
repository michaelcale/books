package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/kjk/u"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func fileForURI(uri string) string {
	path := filepath.Join("www", uri)
	if fileExists(path) {
		return path
	}
	path = path + ".html"
	if fileExists(path) {
		return path
	}
	return ""
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	uri := r.URL.Path
	fmt.Printf("%s\n", uri)
	path := fileForURI(uri)
	if path == "" {
		fmt.Printf("Didn't find file for '%s'\n", uri)
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, path)
}

// https://blog.gopheracademy.com/advent-2016/exposing-go-on-the-internet/
func makeHTTPServer() *http.Server {
	mux := &http.ServeMux{}

	mux.HandleFunc("/", handleIndex)

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second, // introduced in Go 1.8
		Handler:      mux,
	}
	return srv
}

func startPreview() {
	httpSrv := makeHTTPServer()
	httpSrv.Addr = ":8080"

	go func() {
		err := httpSrv.ListenAndServe()
		// mute error caused by Shutdown()
		if err == http.ErrServerClosed {
			err = nil
		}
		u.PanicIfErr(err)
		fmt.Printf("HTTP server shutdown gracefully\n")
	}()
	fmt.Printf("Started listening on %s\n", httpSrv.Addr)
	openBrowser("http://localhost:8080")

	go rebuildOnChanges()

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt /* SIGINT */, syscall.SIGTERM)
	sig := <-c
	fmt.Printf("Got signal %s\n", sig)
}
