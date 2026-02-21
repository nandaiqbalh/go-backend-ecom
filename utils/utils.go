// Package utils offers small helper functions used throughout the server
// code, such as JSON parsing and response writing. Placing them in a
// separate package avoids duplication.
package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// Validate is an instance of the validator used to check struct tags.
var Validate = validator.New()

// ParseJson reads JSON from the request body and unmarshals it into the
// given payload structure. It returns an error if the body is empty or if
// decoding fails.
func ParseJson(r *http.Request, payload any) error {
    if r.Body == nil {
        return fmt.Errorf("request body is empty")
    }

    return json.NewDecoder(r.Body).Decode(payload)
}

// WriteJson sends a JSON response with the specified status code. The
// payload is encoded and written to the response body.
func WriteJson(w http.ResponseWriter, statusCode int, payload any) error {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    return json.NewEncoder(w).Encode(payload)
}

// WriteError is a convenience wrapper that sends an error message in JSON
// form. It calls WriteJson internally.
func WriteError(w http.ResponseWriter, statusCode int, err error) {
    WriteJson(w, statusCode, map[string]string{"error": err.Error()})
} 