package web

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents a json response containing an error message with the reason.
type ErrorResponse struct {
	Message string `json:"error"`
}

func respond(w http.ResponseWriter, data []byte, statusCode int, contentType string) error {
	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return nil
	}

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)

	if _, err := w.Write(data); err != nil {
		return err
	}

	return nil

}

// RespondText is a utility to create an HTTP response in plain text and send it to the client.
func RespondText(w http.ResponseWriter, text string, statusCode int) error {
	return respond(w, []byte(text), statusCode, "text/plain")
}

// RespondJSON is a utility to convert a Go value to JSON and send it to the client.
func RespondJSON(w http.ResponseWriter, data interface{}, statusCode int) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return respond(w, jsonData, statusCode, "application/json")

}

// RespondError is a utility to create an error response and send it to the client.
func RespondError(w http.ResponseWriter, errMsg string, statusCode int) error {
	errResp := ErrorResponse{errMsg}

	return RespondJSON(w, errResp, statusCode)
}
