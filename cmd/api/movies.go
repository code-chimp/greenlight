package main

import (
	"fmt"
	"net/http"
)

// postMoviesHandler creates a new movie
func (app *application) postMoviesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

// getMovieHandler retrieves a movie by ID
func (app *application) getMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "get movie with ID %d\n", id)
}
