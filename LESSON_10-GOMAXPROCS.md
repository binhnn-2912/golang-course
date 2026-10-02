## Bài 10: GOMAXPROCS và giới hạn CPU

### 1. Tổng quan + Khi nào dùng
- `GOMAXPROCS` giới hạn số OS thread tối đa mà Go scheduler dùng để chạy goroutine song song thật sự trên CPU.
- Quan trọng khi benchmark hiệu năng tác vụ CPU-bound, hoặc khi chạy trong container bị giới hạn CPU (cgroup) mà Go không tự nhận đúng số core khả dụng.

### 2. Định nghĩa + Syntax
- `runtime.NumCPU()` trả số logical CPU của máy; `runtime.GOMAXPROCS(n)` đặt số OS thread tối đa chạy goroutine song song.
```go
numCPU := runtime.NumCPU()
fmt.Println("CPU number: ", numCPU)
runtime.GOMAXPROCS(numCPU)

start := time.Now()
var wg sync.WaitGroup
wg.Add(10)
for range 10 {
	go heavyTask(&wg) // CPU-bound, không có I/O wait
}
wg.Wait()
fmt.Println("Total time:", time.Since(start))
```

### 3. Điểm cốt lõi cần nhớ
- Goroutine ≠ OS thread: hàng ngàn goroutine có thể chạy trên vài OS thread; `GOMAXPROCS` giới hạn số OS thread chạy *song song thật sự*, không giới hạn số goroutine được tạo.
- Từ Go 1.5 trở đi, `GOMAXPROCS` mặc định đã bằng `runtime.NumCPU()` — không cần tự set bằng tay như ví dụ trên (chỉ cần khi muốn đổi giá trị khác default).
- Với tác vụ CPU-bound (như `heavyTask` tính tổng vòng lặp lớn), giảm `GOMAXPROCS` xuống thấp hơn số goroutine khiến các goroutine phải thay phiên dùng CPU, tổng thời gian chạy tăng lên.
- `runtime.NumCPU()` trả số CPU logic của máy host, không phải số CPU được cấp phát thực sự khi chạy trong container bị giới hạn (vd Docker `--cpus=2`) — cần thư viện riêng (`go.uber.org/automaxprocs`) để đọc đúng giới hạn cgroup.
- Đo thời gian bằng `time.Since(start)` là cách thực nghiệm đơn giản để so sánh ảnh hưởng của `GOMAXPROCS` lên tác vụ CPU-bound.

### 4. Tránh nhầm lẫn
- `GOMAXPROCS` không giới hạn số goroutine có thể tạo (vẫn tạo được hàng triệu), chỉ giới hạn số goroutine được thực thi song song thật sự tại 1 thời điểm.
- Tăng `GOMAXPROCS` không giúp ích cho tác vụ I/O-bound (chờ mạng, disk) — chỉ có tác dụng rõ rệt với tác vụ CPU-bound như `heavyTask`.
