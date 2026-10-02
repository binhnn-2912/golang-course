## Bài 8: Context

### 1. Tổng quan + Khi nào dùng
- `context.Context` mang tín hiệu "hết hạn/hủy" (deadline, cancel) và truyền xuyên suốt qua các hàm/goroutine con.
- Dùng để kiểm soát thời gian chờ tác vụ bất đồng bộ (gọi API ngoài, query DB...); trong Gin còn dùng để propagate việc hủy khi client đóng kết nối HTTP giữa chừng.

### 2. Định nghĩa + Syntax
- Goroutine thường: tạo `ctx` bằng `context.WithTimeout`, goroutine con dùng `select` lắng nghe `ctx.Done()` song song với channel kết quả.
```go
ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
defer cancel()
go classifyTicket(ctx, ticketID, resultChan)

select {
case result := <-resultChan:
	fmt.Println("Thành công:", result)
case <-ctx.Done():
	err := ctx.Err() // context.DeadlineExceeded hoặc context.Canceled
}

func classifyTicket(ctx context.Context, ticketID int, resultChan chan<- string) {
	select {
	case <-time.After(delay):
		resultChan <- fmt.Sprintf("Ticket %d -> Category: Billing", ticketID)
	case <-ctx.Done():
		return // dừng ngay, không gửi kết quả nữa
	}
}
```
- Trong Gin: mỗi request có sẵn `ctx := c.Request.Context()` — tự hủy khi client đóng kết nối, không cần tự tạo `WithTimeout` riêng cho tầng HTTP.
```go
func (uc *UserController) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()
	result, err := uc.userService.GetUsersService(ctx) // truyền ctx xuống service/repository
}
```

### 3. Điểm cốt lõi cần nhớ
- `ctx.Done()` trả về channel đóng lại khi context hết hạn hoặc bị cancel — dùng trong `select` để biết "nên dừng".
- `ctx.Err()` cho biết lý do dừng: `context.DeadlineExceeded` (hết thời gian) hay `context.Canceled` (bị hủy chủ động) — 2 nguyên nhân cần xử lý nghiệp vụ khác nhau.
- Luôn `defer cancel()` ngay sau khi tạo context bằng `WithTimeout`/`WithCancel` để giải phóng resource dù tác vụ xong sớm hay không.
- `context.Context` nên là tham số đầu tiên của hàm (quy ước Go: `ctx context.Context`), truyền xuyên suốt handler → service → repository, không tự tạo `context.Background()` mới ở tầng giữa.
- Trong Gin, nếu không truyền `c.Request.Context()` xuống các lệnh gọi DB/API, request đã bị client hủy vẫn chạy tiếp ngầm, lãng phí tài nguyên.

### 4. Tránh nhầm lẫn
- `ctx.Done()` không tự động dừng code đang chạy (vd `time.Sleep` hay 1 lệnh gọi không hỗ trợ context) — phải tự check `ctx.Done()`/`ctx.Err()` ở điểm dừng được; context chỉ là "tín hiệu", không phải "ngắt cưỡng bức".
- `gin.Context` (struct request/response của Gin) khác hoàn toàn `context.Context` (chuẩn thư viện) dù tên gần giống — `gin.Context` có field `Request *http.Request`, phải lấy qua `c.Request.Context()` mới ra `context.Context` chuẩn.

### 5. Bài tập
Bạn có 1 hàm giả lập gọi API phân loại ticket (có độ trễ ngẫu nhiên, đôi khi nhanh, đôi khi rất chậm — giống gọi 1 service bên ngoài thật). Bạn cần: nếu API trả lời trong vòng 1 giây thì dùng kết quả đó, nếu quá 1 giây thì bỏ qua, coi như timeout, không đợi mãi.
