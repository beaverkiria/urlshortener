package main

import (
	"github.com/beaverkiria/urlshortener/internal/app/handlers/redirector"
	"github.com/beaverkiria/urlshortener/internal/app/handlers/shortener"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{id}", redirector.Redirect)
	mux.HandleFunc("POST /", shortener.ShortenLink)
	return http.ListenAndServe(`:8080`, mux)
}
