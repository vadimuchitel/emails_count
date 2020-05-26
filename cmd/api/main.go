package main

import (
	"log"
	"net/http"

	"github.com/vadimuchitel/emails_count/internal/middleware"
	"github.com/vadimuchitel/emails_count/internal/processing"
)

func main() {
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(processing.ProcessEmails)
	mux.Handle("/", middleware.ServeHTTP(finalHandler))
	log.Fatal(http.ListenAndServe(":8001", mux))
}
