## Bài 3: Error handler

### 1. Tổng quan + Khi nào dùng
- Dùng 1 envelope response chung (`ResponseData`) và bộ mã lỗi nghiệp vụ (business error code) riêng biệt với HTTP status, để client luôn parse response theo cấu trúc cố định.
- Áp dụng khi API cần phân biệt lỗi tầng nghiệp vụ (không tìm thấy user, sai input...) khỏi lỗi tầng HTTP, để FE xử lý lỗi dựa vào `code` thay vì HTTP status.

### 2. Định nghĩa + Syntax
- `ResponseData` là struct JSON chung cho mọi response; `SuccessResponse`/`ErrorResponse` là helper build response từ code có sẵn trong `msg` map.
```go
// httpStatusCode.go — mã lỗi nghiệp vụ + message tương ứng
const (
	ErrorCodeSuccess    = 20001
	ErrorCodeNotFound   = 20004
)
var msg = map[int]string{ErrorCodeSuccess: "Success", ...}

// response.go — envelope + helper
type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
func SuccessResponse(c *gin.Context, code int, data any) {
	c.JSON(http.StatusOK, ResponseData{Code: code, Message: msg[code], Data: data})
}

// user.controller.go — dùng trong controller
response.SuccessResponse(c, 20001, gin.H{"message": ..., "users": ...})
```

### 3. Điểm cốt lõi cần nhớ
- Mã lỗi nghiệp vụ (`20001`, `20002`...) tách biệt hoàn toàn với HTTP status code — `SuccessResponse` luôn trả `http.StatusOK` (200) dù `code` là gì.
- `msg[code]` tự map code sang message người đọc được, tránh hardcode string message rải rác ở controller.
- `ResponseData` đảm bảo mọi response (thành công hay lỗi) cùng shape `{code, message, data}` — dễ cho FE xử lý đồng nhất.
- `Data` khai báo kiểu `interface{}` (tương đương `any`) để chứa được bất kỳ kiểu dữ liệu nào.

### 4. Tránh nhầm lẫn
- `ErrorResponse` hiện tại vẫn trả `http.StatusOK` (200) giống `SuccessResponse` — dễ nhầm tưởng sẽ trả 4xx/5xx tương ứng, nhưng HTTP status không đổi, chỉ `code` trong body thay đổi.
- Hàm `JSON()` dùng `http.ResponseWriter` thô (không qua Gin) là helper cũ, khác cơ chế với `SuccessResponse`/`ErrorResponse` (dùng `gin.Context`) — không nên dùng lẫn 2 kiểu trong cùng 1 handler.
