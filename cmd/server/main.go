package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"golang-course-api/internal/router"
)

func task(id int, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d bắt đầu \n", id)
	time.Sleep(1 * time.Second)
	ch <- fmt.Sprintf("Task %d hoàn thành", id)
	ch <- fmt.Sprintf("Task %d end", id)
}

func main() {
	r := router.New()

	start := time.Now()

	var wg sync.WaitGroup
	ch := make(chan string, 5)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go task(i, ch, &wg)
	}

	go func() {
		wg.Wait()
		fmt.Println("Đã đóng channel và wait group")
		close(ch)
	}()

	for value := range ch {
		fmt.Println("in ra", value)
	}

	fmt.Printf("Total time: %s\n", time.Since(start))

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
