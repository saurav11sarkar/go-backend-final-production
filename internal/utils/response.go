package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}

func JSON(w http.ResponseWriter, status int, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(Response{Success: status < 400, Message: message, Data: data})
}

func JSONError(w http.ResponseWriter, status int, code, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(Response{
		Success: false,
		Message: message,
		Error:   map[string]string{"code": code, "message": message},
	})
}

// Error is kept as a simple helper for beginner-friendly handlers.
// For errors coming from deeper layers, prefer HandleError.
func Error(w http.ResponseWriter, status int, message string) {
	code := "HTTP_ERROR"
	switch status {
	case 400:
		code = "BAD_REQUEST"
	case 401:
		code = "UNAUTHORIZED"
	case 403:
		code = "FORBIDDEN"
	case 404:
		code = "NOT_FOUND"
	case 409:
		code = "CONFLICT"
	case 422:
		code = "VALIDATION_ERROR"
	case 500:
		code = "INTERNAL_SERVER_ERROR"
	}
	JSONError(w, status, code, message)
}

func Decode(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}
