package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	if os.Getenv("JWT_SECRET") == "" {
		log.Fatal("JWT_SECRET is not set in environment")
	}

	r := mux.NewRouter()
	r.Use(RequestLoggingMiddleware)
	r.Use(SecurityHeadersMiddleware)

	r.HandleFunc("/login", LoginHandler).Methods("POST")

	addr := ":8080"
	log.Println("starting server on", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
