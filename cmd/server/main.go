package main

import (
	"fmt"
	"log"
	"time"

	"golang-course-api/internal/router"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	// TODO: dùng "for job := range jobs" để liên tục nhận job
	// cho tới khi channel jobs bị đóng và rỗng
	// mỗi job: in ra "Worker X đang xử lý job Y", sleep 500ms giả lập việc nặng
	// rồi gửi job*job vào results
	for job := range jobs {
		fmt.Printf("Worker %d đang xử lý job %d\n", id, job)
		time.Sleep(500 * time.Millisecond)
		results <- job * job
	}
}

func main() {
	r := router.New()

	jobs := make(chan int, 3)
	results := make(chan int, 3)

	// Tạo 3 goroutine (worker) để xử lý jobs
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Gửi 5 jobs vào channel jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	close(jobs)
	for i := 1; i <= 5; i++ {
		fmt.Printf("kết quả %d = %d\n", i, <-results)
	}

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
