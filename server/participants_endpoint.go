package main

import (
	"net/http"

	"github.com/dadikovi/race-timer/server/core"
)

type registerParticipantRequest struct {
	StartNumber int64 `json:"startNumber"`
	GroupId     int64 `json:"groupId"`
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
