package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dadikovi/race-timer/server/core"
)

type registerParticipantRequest struct {
	StartNumber int64 `json:"startNumber"`
	GroupId     int64 `json:"groupId"`
}

func (a *App) registerParticipant(w http.ResponseWriter, r *http.Request) {
	var body, bodyReadError = ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if bodyReadError != nil {
		respondWithError(w, http.StatusBadRequest, bodyReadError.Error())
	}

	var request registerParticipantRequest
	if err := json.Unmarshal(body, &request); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	var race, getRaceErr = core.GetRaceInstance(a.DB)
	if getRaceErr != nil {
		respondWithError(w, http.StatusInternalServerError, getRaceErr.Error())
	}

	var createdParticipant, saveParticipantErr = core.MakeParticipantForGroup(request.StartNumber, race.GetActiveGroup()).Save(a.DB)
	if saveParticipantErr != nil {
		respondWithError(w, http.StatusInternalServerError, getRaceErr.Error())
	}

	var result, resultErr = createdParticipant.ToJson()
	if resultErr != nil {
		respondWithError(w, http.StatusInternalServerError, resultErr.Error())
	}

	respondWithJSON(w, http.StatusOK, result)
}
