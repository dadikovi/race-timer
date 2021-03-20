package main

import (
	"net/http"

	"github.com/dadikovi/race-timer/server/core"
)

func (a *App) startActiveGroup(w http.ResponseWriter, r *http.Request) {

	var race, getRaceErr = core.GetRaceInstance(a.DB)
	if getRaceErr != nil {
		respondWithError(w, http.StatusInternalServerError, getRaceErr.Error())
	}

	if err := race.GetActiveGroup().StartGroup(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, []byte(""))
}
