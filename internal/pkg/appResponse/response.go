package appResponse

import (
	"encoding/json"
	"net/http"
)

// JsonResponse Define a helper function to send JSON responses
type JsonResponse struct {
	Code          int         `json:"code"`
	Message       interface{} `json:"message"`
	Data          interface{} `json:"data,omitempty"`
	OptionalField string      `json:"optional_field,omitempty"`
}
type JsonResponseMethod interface {
	GetResponse(w http.ResponseWriter)
}

// GetResponse get response method
func (j JsonResponse) GetResponse(w http.ResponseWriter) {
	err := json.NewEncoder(w).Encode(j)
	if err != nil {
		ResponseMessage(w, http.StatusInternalServerError, "Error encoding response")
		return
	}

}

// ResponseMessage is a helper function that takes in a status code and a message and returns a JSON response
func ResponseMessage(w http.ResponseWriter, status int, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	j := JsonResponse{
		Code:    status,
		Message: message,
	}
	j.GetResponse(w)
	return

}
