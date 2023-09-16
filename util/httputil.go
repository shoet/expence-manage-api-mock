package util

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	ErrorMessageInternalServerError = "InternalServerError"
	ErrorMessageBadRequest          = "BadRequest"
	ErrorMessageNotFound            = "NotFound"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func RespondJSON(w http.ResponseWriter, statusCode int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	jb, err := SerializeToJSON(body)
	if err != nil {
		log.Printf("failed SerializeToJSON: %v", err)
	}

	written, err := w.Write(jb)
	if err != nil {
		log.Printf("failed RespondJSON: %v", err)
	}
	if written == 0 {
		log.Printf("wrtten response is empty")
	}
}

func DeserializeFromJSON(path string, dest any) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed ReadFile %q: %v", path, err)
	}
	if err := json.Unmarshal(b, &dest); err != nil {
		return fmt.Errorf("failed JSON unmarshal %q: %v", path, err)
	}
	return nil
}

func SerializeToJSON(v any) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("failed JSON marshal: %v", err)
	}
	return b, nil
}
