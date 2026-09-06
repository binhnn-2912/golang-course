package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"golang-course-api/internal/router"
)

func printLetters(name string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(delay)
	}
}

func main() {
	r := router.New()

	var wg sync.WaitGroup
	wg.Add(2) // Add 2 goroutines to the WaitGroup
	go printLetters("A", 100*time.Millisecond, &wg)
	go printLetters("B", 150*time.Millisecond, &wg)
	wg.Wait() // Wait for all goroutines to finish

	fmt.Println(("Done"))

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
