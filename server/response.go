package server
import (
	"net/http"
	"encoding/json"
	"log"
)

func SendResponse(w http.ResponseWriter, body interface{}) {
	// Todo check how to set Content type And add a status here
	// 		if there is an HTTPErrorResponse then send 4xx or 5xx error
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Println("Error reporting response", body)
	}
}

func SendErrorResponse(w http.ResponseWriter, reason string) {
	SendResponse(w, HTTPErrorResponse{
		Err: true,
		Reason: reason,
	})
}

func ParseRequestBody(w http.ResponseWriter, r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&v); err != nil {
		log.Println("Could not read response")
		return err
	}
	return nil
}
