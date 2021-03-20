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

	race, err := core.GetRaceInstance(a.DB)
	respondWithServerError(w, err)

	_, err = race.SetActiveGroup(a.DB, savedGroup)
	respondWithServerError(w, err)

	respondWithDto(w, savedGroup.Dto())
}
