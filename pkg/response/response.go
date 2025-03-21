package response

import (
	"encoding/json"
	"net/http"
)

// Response adalah struktur standar untuk response API
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// JSON mengembalikan response JSON dengan format yang konsisten
func JSON(w http.ResponseWriter, status int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

// Error mengembalikan response JSON khusus untuk error
func Error(w http.ResponseWriter, status int, message string) {
	JSON(w, status, message, nil)
}
