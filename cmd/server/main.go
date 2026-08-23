package main

import (
	"log"

	"golang-course-api/internal/router"
)

func main() {
	r := router.New()

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
