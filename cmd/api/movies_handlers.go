package main

import (
	"fmt"
	"github.com/code-chimp/greenlight/internal/models"
	"net/http"
	"time"
)

// getMoviesHandler retrieves a list of movies
func (app *application) getMoviesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "get all movies")
}

// postMoviesHandler creates a new movie
func (app *application) postMoviesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

// getMovieHandler retrieves a movie by ID
func (app *application) getMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	movie := models.Movie{
		ID:             id,
		Title:          "Casablanca",
		Year:           1943,
		RuntimeMinutes: 102,
		Genres:         []string{"drama", "romance", "war"},
		Version:        1,
		CreatedAt:      time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// updateMovieHandler updates a movie by ID
func (app *application) updateMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	fmt.Fprintf(w, "update movie with ID %d\n", id)
}

// deleteMovieHandler deletes a movie by ID
func (app *application) deleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	fmt.Fprintf(w, "delete movie with ID %d\n", id)
}
