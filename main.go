package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/ping", basiHandler)

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("failed to listen ro serve %v", err.Error())
	}
}

func basiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PING!!!"))
}
