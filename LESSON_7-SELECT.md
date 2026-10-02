## Bài 7: Select

### 1. Tổng quan + Khi nào dùng
- `select` cho phép 1 goroutine chờ đồng thời trên nhiều channel, chạy case nào sẵn sàng trước.
- Dùng khi cần đặt timeout cho 1 tác vụ bất đồng bộ (gọi API ngoài, độ trễ không ổn định) mà không muốn chờ vô thời hạn.

### 2. Định nghĩa + Syntax
- Kết hợp `select` với `time.After(d)` — channel trả về 1 giá trị sau khoảng thời gian `d`, dùng làm case "timeout".
```go
func classifyTicket(ticketID int, resultChan chan<- string) {
	delay := time.Duration(200+rand.Intn(1800)) * time.Millisecond
	time.Sleep(delay)
	resultChan <- fmt.Sprintf("Ticket %d -> Category: Billing (mất %v)", ticketID, delay)
}

resultChan := make(chan string)
go classifyTicket(101, resultChan)

select {
case res := <-resultChan:
	fmt.Println("Kết quả dưới 1s: ", res)
case <-time.After(1 * time.Second):
	fmt.Println("Timeout! Dùng category mặc định: Unclassified")
}
```

### 3. Điểm cốt lõi cần nhớ
- `select` block tới khi 1 trong các case có dữ liệu sẵn sàng; nếu nhiều case cùng sẵn sàng, Go chọn ngẫu nhiên 1 case.
- `time.After(d)` trả về channel nhận được giá trị đúng sau khoảng `d` — dùng làm nhánh "hết giờ" trong `select`.
- Goroutine `classifyTicket` vẫn tiếp tục chạy ngầm dù `select` đã chọn nhánh timeout — `select` chỉ ngừng chờ, không hủy goroutine.
- `resultChan` là unbuffered channel — gửi (`resultChan <-`) sẽ block tới khi có người nhận; nếu nhánh timeout được chọn và không còn ai đọc `resultChan` nữa, goroutine `classifyTicket` sẽ bị kẹt mãi ở dòng gửi (goroutine leak).

### 4. Tránh nhầm lẫn
- `select` không phải vòng lặp — chỉ chạy 1 lần; muốn chờ lặp lại nhiều lần phải đặt trong `for { select {...} }`.
- Chọn nhánh timeout không có nghĩa hủy được tác vụ đang chạy — Go không tự động cancel goroutine, muốn hủy thật sự phải dùng `context.WithTimeout` kết hợp channel `Done()`.

### 5. Bài tập
Bạn có 1 hàm giả lập gọi API phân loại ticket (có độ trễ ngẫu nhiên, đôi khi nhanh, đôi khi rất chậm — giống gọi 1 service bên ngoài thật). Bạn cần: nếu API trả lời trong vòng 1 giây thì dùng kết quả đó, nếu quá 1 giây thì bỏ qua, coi như timeout, không đợi mãi.
