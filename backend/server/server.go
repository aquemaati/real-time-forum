package server

import (
	"log"
	"net/http"
	"time"
)

// Function to init a simple server
func InitServer(PORT string, handler http.Handler) *http.Server {

	server := &http.Server{
		Addr: ":" + PORT,

		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return server
}

// Function to run the server
func RunServer(server *http.Server) {
	log.Printf("Server listening on http://localhost%s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
