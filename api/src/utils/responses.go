package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON - return request to JSON format
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// AppError return error to JSON format
func AppError(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Message string `json:"message"`
	}{
		Message: err.Error(),
	})
}
