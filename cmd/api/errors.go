package main

import "net/http"

// logError logs an error with the HTTP request method and URL.
//
// Parameters:
//
//	r - The HTTP request that caused the error.
//	err - The error to log.
func (app *application) logError(r *http.Request, err error) {
	var (
		method = r.Method
		url    = r.URL.RequestURI()
	)
	app.logger.Error(err.Error(), "method", method, "url", url)
}

// errorResponse sends a JSON error response with the specified status code and message.
//
// Parameters:
//
//	w - The HTTP response writer.
//	r - The HTTP request that caused the error.
//	status - The HTTP status code to set in the response.
//	message - The error message to include in the response body.
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// serverErrorResponse logs the error and sends a 500 Internal Server Error response.
//
//	w - The HTTP response writer.
//	r - The HTTP request that caused the error.
//	err - The error to log and include in the response.
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	app.errorResponse(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

// notFoundResponse sends a 404 Not Found response.
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
}

// methodNotAllowedResponse sends a 405 Method Not Allowed response.
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
}
