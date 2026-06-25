package server

import (
	"net/http"
)

func NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}

func Start() error {
	mux := NewServeMux()

	newServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return newServer.ListenAndServe()
}
