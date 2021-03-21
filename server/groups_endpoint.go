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
	parseRequestBody(w, r, &request)

	s, err := core.FetchSegmentById(a.DB, request.SegmentId)
	respondWithClientError(w, err)

	var group = core.MakeGroupForSegment(s)
	savedGroup, err := group.Save(a.DB)
	respondWithClientError(w, err)

	_, err = a.race.SetActiveGroup(a.DB, savedGroup)
	respondWithServerError(w, err)

	err = a.RefreshRace()
	respondWithServerError(w, err)

	respondWithDto(w, savedGroup.Dto())
}

func (a *App) startActiveGroup(w http.ResponseWriter, r *http.Request) {

	g, err := a.race.GetActiveGroup().StartGroup(a.DB)
	respondWithServerError(w, err)

	respondWithDto(w, g.Dto())
}
