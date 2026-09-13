## Bài 5: Channel (buffered / unbuffered) — Worker Pool

### 1. Tổng quan + Khi nào dùng
- Channel là cơ chế giao tiếp an toàn giữa các goroutine, built-in trong Go, tránh phải tự quản lý lock/mutex khi truyền dữ liệu.
- Worker pool dùng channel để giới hạn số goroutine xử lý đồng thời (vd 3 worker cố định), thay vì tạo goroutine riêng cho từng job.

### 2. Định nghĩa + Syntax
- Unbuffered channel (`make(chan T)`): gửi block tới khi có người nhận. Buffered channel (`make(chan T, n)`): gửi không block cho tới khi buffer đầy.
```go
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs { // nhận job tới khi channel đóng và rỗng
		fmt.Printf("Worker %d đang xử lý job %d\n", id, job)
		time.Sleep(500 * time.Millisecond)
		results <- job * job
	}
}

jobs := make(chan int, 3)    // buffered, sức chứa 3
results := make(chan int, 3)

for w := 1; w <= 3; w++ {
	go worker(w, jobs, results) // 3 worker cố định
}
for j := 1; j <= 5; j++ {
	jobs <- j
}
close(jobs)
for i := 1; i <= 5; i++ {
	fmt.Printf("kết quả %d = %d\n", i, <-results)
}
```

### 3. Điểm cốt lõi cần nhớ
- `jobs <-chan int` (receive-only) và `results chan<- int` (send-only) là directional channel — giới hạn worker chỉ đọc `jobs`, chỉ ghi `results`, compiler check lúc biên dịch.
- `for job := range jobs` tự dừng khi channel bị `close()` và đã rỗng, không cần điều kiện thoát thủ công.
- Phải `close(jobs)` sau khi gửi đủ job — nếu không, `range` sẽ block mãi chờ job tiếp theo.
- Buffered channel dung lượng 3 cho phép gửi tối đa 3 giá trị mà chưa cần người nhận ngay; gửi giá trị thứ 4 sẽ block tới khi có người lấy bớt ra.
- Số worker (3) độc lập với số job (5) — worker rảnh tự lấy job tiếp theo từ channel, đúng ý nghĩa "pool".

### 4. Tránh nhầm lẫn
- Buffered channel không có nghĩa "không block" — chỉ block trễ hơn (khi buffer đầy), không loại bỏ hoàn toàn tính đồng bộ.
- `close(jobs)` không xóa dữ liệu còn lại trong buffer — receiver vẫn đọc hết các giá trị đã gửi trước đó, chỉ là không gửi thêm được nữa.
- Không cần `close(results)` trong bài này vì `main` đã biết trước số lượng kết quả cần đọc (5); chỉ bắt buộc đóng channel khi phía nhận dùng `range` để biết lúc nào dừng.

### 5. Bài tập
Worker Pool xử lý công việc

Bối cảnh: Bạn có 5 "job" (số nguyên từ 1-5) cần xử lý (giả lập bằng cách bình phương số đó). Thay vì tạo 5 goroutine riêng lẻ, bạn dùng 3 worker cố định để xử lý — giống connection pool hay thread pool.
Yêu cầu:
Tạo channel jobs để gửi công việc vào
Tạo channel results để nhận kết quả ra
Tạo 3 worker (goroutine), mỗi worker liên tục lấy job từ jobs, xử lý, gửi kết quả vào results
main gửi 5 job vào, đóng channel jobs, rồi thu thập đủ 5 kết quả
