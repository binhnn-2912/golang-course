## Bài 9: Race condition + Mutex

### 1. Tổng quan + Khi nào dùng
- Race condition xảy ra khi nhiều goroutine cùng đọc/ghi 1 biến dùng chung mà không đồng bộ hóa, cho kết quả không xác định (đôi khi đúng, đôi khi sai, tùy lần chạy).
- `sync.Mutex` dùng để khóa 1 đoạn code (critical section) sao cho chỉ 1 goroutine thực thi tại 1 thời điểm, tránh race condition khi nhiều goroutine cùng sửa 1 biến/struct chung.

### 2. Định nghĩa + Syntax
- `mu.Lock()` trước và `mu.Unlock()` sau đoạn code cần bảo vệ — đảm bảo chỉ 1 goroutine vào được đoạn đó tại 1 thời điểm.
```go
counter := 0
var wg sync.WaitGroup
var mu sync.Mutex

for i := 0; i < 1000; i++ {
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		counter++
		mu.Unlock()
	}()
}

wg.Wait()
fmt.Println("Counter cuối cùng:", counter)
```

### 3. Điểm cốt lõi cần nhớ
- `counter++` thực chất là 3 bước (đọc - cộng - ghi), không phải 1 phép toán nguyên tử — 2 goroutine cùng đọc giá trị cũ rồi ghi đè nhau là nguyên nhân gây race condition.
- `mu.Lock()`/`mu.Unlock()` đảm bảo đoạn `counter++` chạy tuần tự (serialize) giữa các goroutine, dù các goroutine vẫn chạy song song ở phần còn lại.
- Luôn `Unlock()` sau khi `Lock()` — nên dùng `defer mu.Unlock()` ngay sau dòng `Lock()` để tránh quên mở khóa khi có nhiều nhánh return/panic.
- Critical section (vùng giữa `Lock()`/`Unlock()`) nên càng ngắn càng tốt — khóa lâu làm các goroutine khác phải chờ, giảm hiệu năng.
- `go run -race` (race detector) giúp phát hiện race condition lúc dev/test, vì bug này không phải lúc nào cũng lộ ra khi chạy thường.

### 4. Tránh nhầm lẫn
- Mutex không phải channel — Mutex bảo vệ *trạng thái dùng chung* (shared state), channel dùng để *giao tiếp/truyền dữ liệu* giữa goroutine; chọn nhầm công cụ khiến code phức tạp hơn cần thiết.
- Quên `Lock()` không gây lỗi biên dịch hay panic — chương trình vẫn chạy, chỉ cho kết quả sai (counter < 1000) một cách ngẫu nhiên, rất khó phát hiện nếu không test kỹ hoặc dùng `-race`.
