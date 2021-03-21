package main

import (
	"net/http"
	"strconv"

	"github.com/dadikovi/race-timer/server/core"
	"github.com/gorilla/mux"
)

type registerParticipantRequest struct {
	StartNumber int `json:"startNumber"`
	GroupId     int `json:"groupId"`
}

func (a *App) registerParticipant(w http.ResponseWriter, r *http.Request) {

	var request registerParticipantRequest
	parseRequestBody(w, r, &request)

	race, err := core.GetRaceInstance(a.DB)
	respondWithServerError(w, err)

	createdParticipant, err := core.MakeParticipantForGroup(request.StartNumber, race.GetActiveGroup()).Save(a.DB)
	respondWithServerError(w, err)

	respondWithDto(w, createdParticipant.Dto())
}

func (a *App) participantFinished(w http.ResponseWriter, r *http.Request) {
	startNumber, err := strconv.ParseInt(mux.Vars(r)["startNumber"], 10, 64)
	respondWithClientError(w, err)

	p, err := core.FetchParticipantByStartNumber(a.DB, int(startNumber))
	respondWithServerError(w, err)

	p, err = p.Finish(a.DB)
	respondWithServerError(w, err)

	respondWithDto(w, p.Dto())
}
