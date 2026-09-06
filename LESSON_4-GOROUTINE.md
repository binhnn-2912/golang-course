## Bài 4: Goroutine

### 1. Tổng quan + Khi nào dùng
- Goroutine là đơn vị thực thi đồng thời (concurrent) nhẹ của Go, do Go runtime tự quản lý và lên lịch, không phải OS thread trực tiếp.
- Dùng khi cần chạy nhiều tác vụ độc lập song song (I/O, tính toán) mà không muốn block luồng chính.

### 2. Định nghĩa + Syntax
- `go funcCall()` khởi chạy 1 goroutine mới; `sync.WaitGroup` dùng để đợi các goroutine hoàn thành trước khi tiếp tục.
```go
func printLetters(name string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 3; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(delay)
	}
}

var wg sync.WaitGroup
wg.Add(2) // khai báo số goroutine cần đợi
go printLetters("A", 100*time.Millisecond, &wg)
go printLetters("B", 150*time.Millisecond, &wg)
wg.Wait() // block tới khi cả 2 goroutine gọi Done()
```

### 3. Điểm cốt lõi cần nhớ
- Từ khóa `go` đặt trước lời gọi hàm sẽ chạy hàm đó trong goroutine mới, không block code phía sau.
- `wg.Add(n)` khai báo số lượng goroutine cần đợi, gọi trước khi các goroutine bắt đầu chạy.
- `defer wg.Done()` đảm bảo giảm counter dù hàm return bình thường hay panic giữa chừng.
- `wg.Wait()` block cho tới khi counter về 0 — tránh `main` thoát sớm khi goroutine chưa chạy xong.
- Thứ tự in ra giữa các goroutine không xác định (interleaved), phụ thuộc scheduler và độ trễ riêng của từng goroutine.

### 4. Tránh nhầm lẫn
- Goroutine không phải OS thread — nhẹ hơn nhiều, được Go runtime lên lịch (M:N scheduling) trên số lượng OS thread giới hạn.
- Gọi `Done()` nhiều hơn số lần `Add()` gây panic "negative WaitGroup counter"; quên `wg.Wait()` khiến `main` có thể thoát trước khi goroutine chạy xong.
