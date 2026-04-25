package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// JSONResponse defines the standard format for all JSON responses.
type JSONResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// ReadJSON reads JSON from the request body and decodes it into data.
// The data parameter must be a pointer. It enforces a maximum body size
// and ensures that only a single JSON object is present in the request.
// It returns an error if the body is empty, too large, or malformed.
func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1048576 // 1 Megabyte

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)

	dec.DisallowUnknownFields()

	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})

	if err != io.EOF {
		return errors.New("body must contain only one JSON object")
	}

	if err == io.EOF {
		return errors.New("empty request body")
	}

	return nil
}

// WriteJSON writes the given data as JSON to the response writer with the
// specified HTTP status code. Optional headers can be provided and will
// be added to the response before writing.
func WriteJSON(w http.ResponseWriter, code int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, values := range headers[0] {
			for _, v := range values {
				w.Header().Add(key, v)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

// ErrorJSON writes a JSON error response with the given error message.
// An optional HTTP status code can be provided; otherwise, it defaults
// to http.StatusBadRequest.
func ErrorJSON(w http.ResponseWriter, err error, code ...int) error {
	statusCode := http.StatusBadRequest

	if len(code) > 0 {
		statusCode = code[0]
	}

	var payload JSONResponse
	payload.Error = true
	payload.Message = err.Error()

	return WriteJSON(w, statusCode, payload)
}
