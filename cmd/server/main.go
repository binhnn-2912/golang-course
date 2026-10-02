package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func heavyTask(wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for i := 0; i < 100e8; i++ {
		sum++
	}
	fmt.Println("Sum:", sum)
}

func main() {
	// r := router.New()

	numCPU := runtime.NumCPU()
	fmt.Println("CPU number: ", numCPU)

	runtime.GOMAXPROCS(numCPU)

	start := time.Now()

	var wg sync.WaitGroup

	wg.Add(10)

	for range 10 {
		go heavyTask(&wg)
	}

	wg.Wait()
	fmt.Println("Total time:", time.Since(start))

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	// if err := r.Run(); err != nil {
	// 	log.Fatalf("failed to run server: %v", err)
	// }
}
