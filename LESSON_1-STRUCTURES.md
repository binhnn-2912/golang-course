## Bài 1: Cấu trúc của 1 dự án Go

### 1. Tổng quan + Khi nào dùng
- Go không ép cấu trúc thư mục cố định, nhưng cộng đồng dùng chung "Standard Go Project Layout" để tách rõ: entry point, logic nội bộ, code dùng chung, config.
- Áp dụng cho service/API vừa-lớn, nhiều người tham gia; script nhỏ thì không cần.

### 2. Định nghĩa + Syntax
- `cmd/` = entry point (mỗi sub-folder = 1 binary). `internal/` = code riêng, chia layer (controller → service → repository → domain, cộng thêm router, config). `pkg/` = code dùng chung, export được cho project khác.
```go
// cmd/server/main.go — entry point
func main() {
	r := router.New()
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

// internal/router/router.go — nối các layer
// router -> controller -> service -> repository -> models -> db
func New() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	v1.GET("/users", controller.NewUserController().GetUsers)
	return r
}
```

### 3. Điểm cốt lõi cần nhớ
- Module name trong `go.mod` (`golang-course-api`) là prefix cho mọi import nội bộ (`golang-course-api/internal/...`).
- Luồng chạy: `main.go` → `router` → `controller` → `service` → `repository` → `domain`.
- `internal/`: Go compiler chặn import từ ngoài module — chỉ code cùng module mới dùng được.
- `pkg/`: ngược lại `internal`, thiết kế để chia sẻ/tái dùng ở project khác.
- Mỗi thư mục = 1 package; `package main` bắt buộc để có hàm `main()` chạy được.
- `configs/`, `migrations/`, `scripts/`, `tests/`, `third_party/` hiện trống — chỗ dự phòng theo convention.

### 4. Tránh nhầm lẫn
- `internal/` vs `pkg/`: internal giới hạn trong module, pkg chia sẻ ra ngoài.
- `go.mod` khai báo module (toàn bộ project), mỗi thư mục con mới là package.
- Thư mục `response/` ở gốc (trống) khác `pkg/response/` (có code) — dễ nhầm vì trùng tên.
