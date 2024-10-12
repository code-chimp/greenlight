package main

import (
	"github.com/code-chimp/greenlight/internal/models"
	"net/http"
)

// healthcheckHandler returns the status of the application
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := models.SystemInfo{
		Environment: app.config.env,
		Version:     version,
		Revision:    revision,
		Status:      "available",
	}

	err := app.writeJSON(w, http.StatusOK, envelope{"system_info": data}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
