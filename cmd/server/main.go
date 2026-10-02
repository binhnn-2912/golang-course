package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"golang-course-api/internal/router"
)

func classifyTicket(ticketID int, resultChan chan<- string) {
	delay := time.Duration(200+rand.Intn(1800)) * time.Millisecond
	log.Println("Đang xử lý ticket ", ticketID, "tốn ", delay)
	time.Sleep(delay)
	resultChan <- fmt.Sprintf("Ticket %d -> Category: Billing (mất %v)", ticketID, delay)
}

func main() {
	r := router.New()

	ticketID := 101
	resultChan := make(chan string)

	go classifyTicket(ticketID, resultChan)

	// TODO: dùng select với 2 case:
	// case 1: nhận được kết quả từ resultChan -> in ra kết quả
	// case 2: sau 1 giây (time.After) -> in ra "Timeout! Dùng category mặc định: Unclassified"

	select {
	case res := <-resultChan:
		fmt.Println("Kết quả dưới 1s: ", res)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout! Dùng category mặc định: Unclassified")
	}

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
