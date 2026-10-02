package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"golang-course-api/internal/router"
)

func classifyTicket(ctx context.Context, ticketID int, resultChan chan<- string) {
	delay := time.Duration(200+rand.Intn(1800)) * time.Millisecond
	fmt.Println("Delay time: ", delay)
	select {
	case <-time.After(delay):
		// giả lập API trả về kết quả sau "delay"
		resultChan <- fmt.Sprintf("Ticket %d -> Category: Billing (mất %v)", ticketID, delay)
	case <-ctx.Done():
		// TODO: context bị hủy (timeout hoặc cancel) trước khi kịp xong
		// in ra: "Ticket X bị hủy: " + lý do (ctx.Err())
		// return luôn, không gửi gì vào resultChan nữa
		return
	}
}

func main() {
	r := router.New()

	ticketID := 101
	resultChan := make(chan string)

	// TODO: tạo ctx với context.WithTimeout, timeout = 1 giây
	// TODO: nhớ defer cancel() để giải phóng resource

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go classifyTicket(ctx, ticketID, resultChan)

	// TODO: dùng select nhận kết quả từ resultChan HOẶC ctx.Done()
	// (gợi ý: lúc này 2 nơi đều có thể biết "hết giờ" — main và goroutine
	// đều lắng nghe cùng 1 ctx.Done(), không cần time.After riêng ở main nữa)

	select {
	case result := <-resultChan:
		// Trường hợp thành công — có kết quả thật từ "API"
		fmt.Println("Thành công:", result)

	case <-ctx.Done():
		// Trường hợp bị hủy — cần phân biệt RÕ 2 nguyên nhân khác nhau
		// để xử lý nghiệp vụ đúng cách (retry/alert khác nhau hoàn toàn)
		err := ctx.Err()

		switch {
		case errors.Is(err, context.DeadlineExceeded):
			// API/hệ thống xử lý quá chậm, vượt quá thời gian cho phép (1s).
			// Đây là vấn đề về PERFORMANCE — nên log lại để theo dõi,
			// và có thể cân nhắc retry hoặc đưa vào hàng đợi xử lý lại.
			fmt.Printf("Ticket %d: TIMEOUT sau 1s — hệ thống phân loại quá chậm. "+
				"Cần log/alert để theo dõi performance. (err: %v)\n", ticketID, err)
			// TODO thực tế: ghi log structured, tăng metric "classify_timeout_total",
			// đưa ticketID vào queue retry nếu cần.

		case errors.Is(err, context.Canceled):
			// Bị hủy chủ động — ví dụ user đóng tab, request HTTP gốc bị ngắt,
			// hoặc 1 phần khác của hệ thống chủ động cancel.
			// Đây KHÔNG phải lỗi hệ thống — không cần retry, không cần alert.
			fmt.Printf("Ticket %d: Bị HỦY chủ động (không phải lỗi hệ thống). "+
				"Không cần retry. (err: %v)\n", ticketID, err)
			// TODO thực tế: có thể chỉ cần log ở mức debug, hoặc bỏ qua hoàn toàn.

		default:
			// Trường hợp hiếm gặp: ctx.Err() trả về lỗi khác (thường không xảy ra
			// với context.WithTimeout/WithCancel chuẩn), nhưng vẫn nên có nhánh
			// dự phòng để không bỏ sót trường hợp lạ.
			fmt.Printf("Ticket %d: Lỗi context không xác định: %v\n", ticketID, err)
		}
	}
	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
