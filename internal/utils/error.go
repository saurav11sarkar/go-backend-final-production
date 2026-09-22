package utils

import "net/http"

// AppError is the standard error returned by the application.
// Keep business errors small and predictable so handlers stay easy to read.
type AppError struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

func NewAppError(status int, code, message string) *AppError {
	return &AppError{Status: status, Code: code, Message: message}
}

func HandleError(w http.ResponseWriter, err error) {
	if appErr, ok := err.(*AppError); ok {
		JSONError(w, appErr.Status, appErr.Code, appErr.Message)
		return
	}
	JSONError(w, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "internal server error")
}
