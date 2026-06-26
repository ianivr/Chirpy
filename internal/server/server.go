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

	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("./"))))
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(http.StatusText(http.StatusOK)))
	})

	newServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return newServer.ListenAndServe()
}
