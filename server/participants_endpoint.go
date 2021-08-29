package main

import (
	"log"
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
	if err := parseRequestBody(w, r, &request); err != nil {
		return
	}

	createdParticipant, err := core.MakeParticipantForGroup(request.StartNumber, a.race.GetActiveGroup()).Save(a.DB)
	respondWithServerError(w, err)
	if err != nil {
		return
	}

	respondWithDto(w, createdParticipant.Dto())
}

func (a *App) participantFinished(w http.ResponseWriter, r *http.Request) {
	startNumber, err := strconv.ParseInt(mux.Vars(r)["startNumber"], 10, 64)
	log.Print(err)
	respondWithClientError(w, err)
	if err != nil {
		return
	}

	p, err := core.FetchParticipantByStartNumber(a.DB, int(startNumber))
	log.Print(err)
	respondWithServerError(w, err)
	if err != nil {
		return
	}

	p, err = p.Finish(a.DB)
	log.Print(err)
	respondWithServerError(w, err)
	if err != nil {
		return
	}

	respondWithDto(w, p.Dto())
}
