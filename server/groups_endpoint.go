package main

import (
	"net/http"

	"github.com/dadikovi/race-timer/server/core"
)

type createGroupRequest struct {
	SegmentId int `json:"segmentId"`
}

func (a *App) createGroup(w http.ResponseWriter, r *http.Request) {

	var request createGroupRequest
	if err := parseRequestBody(w, r, &request); err != nil {
		return
	}

	s, err := core.FetchSegmentById(a.DB, request.SegmentId)
	respondWithClientError(w, err)
	if err != nil {
		return
	}

	var group = core.MakeGroupForSegment(s)
	savedGroup, err := group.Save(a.DB)
	respondWithClientError(w, err)
	if err != nil {
		return
	}

	_, err = a.race.SetActiveGroup(a.DB, savedGroup)
	respondWithServerError(w, err)
	if err != nil {
		return
	}

	err = a.RefreshRace()
	respondWithServerError(w, err)
	if err != nil {
		return
	}

	respondWithDto(w, savedGroup.Dto())
}

func (a *App) startActiveGroup(w http.ResponseWriter, r *http.Request) {

	g, err := a.race.GetActiveGroup().StartGroup(a.DB)
	respondWithServerError(w, err)
	if err != nil {
		return
	}

	respondWithDto(w, g.Dto())
}
