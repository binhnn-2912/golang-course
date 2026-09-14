## Bài 6: WaitGroup + Channel

### 1. Tổng quan + Khi nào dùng
- Kết hợp `WaitGroup` (đợi nhiều goroutine hoàn thành) với `channel` (thu thập kết quả) để nhiều goroutine gửi dữ liệu về 1 nơi tập trung mà `main` không cần biết trước tổng số lượng item.
- Dùng khi số lượng item mỗi goroutine gửi ra không cố định (ở đây mỗi `task` gửi 2 message), nên không thể dùng vòng lặp đếm cứng như bài Worker Pool (Bài 5).

### 2. Định nghĩa + Syntax
- Pattern chuẩn: các goroutine gửi vào channel, 1 goroutine riêng đợi `wg.Wait()` rồi `close(channel)`, còn `main` dùng `range` để nhận tới khi channel đóng.
```go
func task(id int, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	ch <- fmt.Sprintf("Task %d hoàn thành", id)
	ch <- fmt.Sprintf("Task %d end", id)
}

var wg sync.WaitGroup
ch := make(chan string, 5)
for i := 1; i <= 5; i++ {
	wg.Add(1)
	go task(i, ch, &wg)
}

go func() { // goroutine riêng để đóng channel
	wg.Wait()
	close(ch)
}()

for value := range ch { // main chỉ lo nhận, không lo đóng
	fmt.Println("in ra", value)
}
```

### 3. Điểm cốt lõi cần nhớ
- 5 task, mỗi task gửi 2 message → tổng 10 message cần channel xử lý, nhưng buffer chỉ có 5 — nếu không có ai nhận song song, channel sẽ đầy và block.
- Buffered channel không tránh được deadlock nếu số lượng dữ liệu gửi vào vượt quá capacity mà không có receiver đang chạy — Go runtime báo lỗi "all goroutines are asleep - deadlock!".
- Vì vậy `wg.Wait()` + `close(ch)` phải nằm trong 1 goroutine riêng, chạy song song với vòng `for range ch` ở `main` — để receiver luôn sẵn sàng nhận trong lúc các task vẫn đang gửi.
- Nếu viết `wg.Wait(); close(ch)` tuần tự ngay trước `for range ch` trên cùng 1 goroutine, `main` bị block tại `wg.Wait()` chờ task xong, trong khi task lại bị block chờ buffer trống — kẹt lẫn nhau (deadlock).

### 4. Tránh nhầm lẫn
- Quy tắc chung: **luôn tách `close()` ra goroutine riêng khi bên nhận dùng `range`** — áp dụng cho mọi channel, kể cả buffered lẫn unbuffered, không phải chỉ khi thấy deadlock mới cần tách.
- Tăng buffer không phải cách fix đúng cho lỗi đặt `close()` sai chỗ — chỉ che giấu vấn đề tới khi dữ liệu gửi vào nhiều hơn buffer thì lỗi mới lộ ra.
