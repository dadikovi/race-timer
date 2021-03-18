package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dadikovi/race-timer/server/core"
)

type createGroupRequest struct {
	SegmentId int64 `json:"segmentId"`
}

func (a *App) createGroup(w http.ResponseWriter, r *http.Request) {
	var body, bodyReadError = ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if bodyReadError != nil {
		respondWithError(w, http.StatusBadRequest, bodyReadError.Error())
	}

	var request createGroupRequest
	if err := json.Unmarshal(body, &request); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	var s, fetchSegmentErr = core.FetchSegmentById(a.DB, request.SegmentId)
	if fetchSegmentErr != nil {
		respondWithError(w, http.StatusBadRequest, fetchSegmentErr.Error())
	}

	var group = core.MakeGroupForSegment(s)
	var savedGroup, groupSaveErr = group.Save(a.DB)
	if groupSaveErr != nil {
		respondWithError(w, http.StatusBadRequest, groupSaveErr.Error())
	}

	var race, getRaceErr = core.GetRaceInstance(a.DB)
	if getRaceErr != nil {
		respondWithError(w, http.StatusInternalServerError, getRaceErr.Error())
	}

	var _, setActiveGroupErr = race.SetActiveGroup(a.DB, savedGroup)
	if setActiveGroupErr != nil {
		respondWithError(w, http.StatusInternalServerError, setActiveGroupErr.Error())
	}

	var result, resultErr = savedGroup.ToJson()
	if resultErr != nil {
		respondWithError(w, http.StatusInternalServerError, resultErr.Error())
	}

	respondWithJSON(w, http.StatusOK, result)
}
