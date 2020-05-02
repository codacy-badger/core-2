package server

import (
	"log"
	"net/http"
	"os"
)

// CreateServer creates a HTTP server
func CreateServer() {
	port := os.Getenv("IBORG_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Go HTTP Server running on port %v\n", ":"+port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
