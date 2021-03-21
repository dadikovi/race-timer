package main

import (
	"net/http"

	"github.com/dadikovi/race-timer/server/core"
)

func (a *App) startActiveGroup(w http.ResponseWriter, r *http.Request) {

	race, err := core.GetRaceInstance(a.DB)
	respondWithServerError(w, err)

	g, err := race.GetActiveGroup().StartGroup(a.DB)
	respondWithServerError(w, err)

	respondWithDto(w, g.Dto())
}
