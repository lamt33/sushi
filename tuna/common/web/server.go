package web

import (
	"log"
	"net/http"
	"os"
	"time"
)

func Server(handlerFunc http.HandlerFunc) *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to Port %s", port)
	}

	return &http.Server{
		Addr:    ":" + port,
		Handler: handlerFunc,
		//TLSConfig:    nil,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}
