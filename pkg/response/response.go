package response

import (
	"encoding/json"
	"net/http"
)

// Envelope wraps any response body.
type Envelope map[string]any

// JSON writes a JSON response with the given status code.
func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// Success writes a standard success envelope.
func Success(w http.ResponseWriter, status int, data any) {
	JSON(w, status, Envelope{
		"success": true,
		"data":    data,
	})
}

// Error writes a standard error envelope.
func Error(w http.ResponseWriter, status int, message string) {
	JSON(w, status, Envelope{
		"success": false,
		"error":   message,
	})
}

// ValidationError writes a 422 with field-level error details.
func ValidationError(w http.ResponseWriter, errors map[string]string) {
	JSON(w, http.StatusUnprocessableEntity, Envelope{
		"success": false,
		"errors":  errors,
	})
}
