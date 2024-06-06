package utils

import (
	"encoding/json"
	"net/http"
)

type AppError struct {
	Message string `json:"message"`
	Code    int    `json:"-"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(message string, code int) *AppError {
	return &AppError{Message: message, Code: code}
}

func RespondWithError(w http.ResponseWriter, err *AppError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code)
	json.NewEncoder(w).Encode(err)
}
