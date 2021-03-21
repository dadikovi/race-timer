package main

import (
	"net/http"

	"github.com/dadikovi/race-timer/server/core"
)

func (a *App) fetchAllSegment(w http.ResponseWriter, r *http.Request) {

	segments, err := core.FetchAll(a.DB)
	respondWithClientError(w, err)

	var results []core.SegmentDto
	for _, s := range segments {
		results = append(results, s.Dto())
	}

	respondWithDto(w, results)
}

type createSegmentRequest struct {
	Name string `json:"name"`
}

func (a *App) createSegment(w http.ResponseWriter, r *http.Request) {
	var request createSegmentRequest
	parseRequestBody(w, r, &request)

	if s, err := core.MakeSegment(request.Name).Save(a.DB); err != nil {
		if err.Error() == core.ALREADY_EXISTS_ERROR_CODE {
			respondWithClientError(w, err)
		}
		respondWithServerError(w, err)
	} else {
		respondWithDto(w, s.Dto())
	}
}
