## Bài 2: Gin với MVC

### 1. Tổng quan + Khi nào dùng
- Gin là web framework HTTP cho Go, cung cấp router, middleware, `Context` để xử lý request/response nhanh gọn.
- Kết hợp Gin với pattern layer kiểu MVC (Controller - Service - Repository) để tách handler HTTP khỏi logic nghiệp vụ và truy xuất dữ liệu, dễ test và mở rộng.

### 2. Định nghĩa + Syntax
- `gin.Engine` là router chính; `gin.Context` đại diện 1 request/response, dùng để đọc input và trả output.
```go
// router.go — khởi tạo engine, group route theo version, gắn handler
r := gin.Default()
v1 := r.Group("/api/v1")
v1.GET("/users", controller.NewUserController().GetUsers)

// user.controller.go — Controller nhận request, gọi Service, trả JSON
func (uc *UserController) GetUsers(c *gin.Context) {
	name := c.DefaultQuery("name", "Hương") // query param, có default
	uid := c.Query("uid")                   // query param, rỗng nếu thiếu

	c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetUsersService(),
		"name": name, "uid": uid,
	})
}
```

### 3. Điểm cốt lõi cần nhớ
- `gin.Default()` = router có sẵn middleware Logger + Recovery; `gin.New()` nếu muốn router trơn.
- `r.Group(prefix)` tạo route group dùng chung prefix, có thể gắn middleware riêng cho cả group.
- `c.Query(key)` lấy query param (rỗng nếu thiếu); `c.DefaultQuery(key, default)` có giá trị mặc định.
- `gin.H` là alias của `map[string]any`, dùng build JSON response nhanh.
- Controller không gọi thẳng Repository — luôn qua Service (`UserController` → `UserService` → `UserRepository`), giữ đúng chiều phụ thuộc.
- Mỗi layer dùng constructor `NewXxx()` để khởi tạo và inject dependency (Controller inject Service, Service inject Repository).

### 4. Tránh nhầm lẫn
- Gin không phải framework MVC — đây là cách dự án tự tổ chức code theo layer, Gin chỉ lo routing/HTTP.
- Trong `GetUsers`, chỉ field `message` đi qua chuỗi Service → Repository; mảng `users` đang hardcode thẳng trong Controller — dễ nhầm tưởng cả response đều lấy từ Repository.
