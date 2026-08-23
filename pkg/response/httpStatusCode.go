package response

const (
	ErrorCodeSuccess       = 20001 // Success
	ErrorCodeBadRequest    = 20002 // Bad Request
	ErrorCodeUnauthorized  = 20003 // Unauthorized
	ErrorCodeNotFound      = 20004 // Not Found
	ErrorCodeInternalError = 20005 // Internal Error
)

var msg = map[int]string{
	ErrorCodeSuccess:       "Success",
	ErrorCodeBadRequest:    "Bad Request",
	ErrorCodeUnauthorized:  "Unauthorized",
	ErrorCodeNotFound:      "Not Found",
	ErrorCodeInternalError: "Internal Error",
}
