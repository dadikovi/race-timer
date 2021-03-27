package main

import (
	"net/http"
)

func (a *App) getResults(w http.ResponseWriter, r *http.Request) {

	g, err := a.race.Results(a.DB)
	respondWithServerError(w, err)
	if err != nil {
		return
	}

	respondWithDto(w, g)
}
