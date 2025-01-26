package main

import "net/http"

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if err := writeJSON(w, http.StatusOK, "ok"); err != nil {

	}
}
