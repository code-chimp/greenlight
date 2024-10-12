package main

import (
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"maps"
	"net/http"
	"strconv"
)

// readIDParam reads the ID parameter from the request context.
// It expects the ID to be a positive integer.
// If the ID is invalid or less than 1, it returns an error.
//
// Parameters:
//
//	r - The HTTP request containing the ID parameter.
//
// Returns:
//
//	int64 - The parsed ID.
//	error - An error if the ID is invalid.
func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

type envelope map[string]any

// writeJSON writes a JSON response to the http.ResponseWriter.
// It marshals the provided data to JSON, sets the appropriate headers,
// and writes the JSON response with the specified status code.
//
// Parameters:
//
//	w - The HTTP response writer.
//	status - The HTTP status code to set in the response.
//	data - The data to marshal to JSON and include in the response body.
//	headers - Additional headers to include in the response.
//
// Returns:
//
//	error - An error if JSON marshaling fails.
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	js = append(js, '\n')

	maps.Insert(w.Header(), maps.All(headers))

	// Write the JSON response to the http.ResponseWriter
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
