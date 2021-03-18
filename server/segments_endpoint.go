package main

import (
	"io/ioutil"
	"net/http"

	"github.com/dadikovi/race-timer/server/core"
)

func (a *App) fetchAllSegment(w http.ResponseWriter, r *http.Request) {
	var json = "["
	var segments, err = core.FetchAll(a.DB)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	for _, s := range segments {
		sjson, _ := s.ToJson()
		json = json + string(sjson)
	}

	json = json + "]"

	respondWithJSON(w, http.StatusOK, []byte(json))
}

func (a *App) createSegment(w http.ResponseWriter, r *http.Request) {
	var body, bodyReadError = ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if bodyReadError != nil {
		respondWithError(w, http.StatusBadRequest, bodyReadError.Error())
	}

	var s, err = core.MakeSegment(string(body))

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	if err := s.Save(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var json, _ = s.ToJson()
	respondWithJSON(w, http.StatusCreated, json)
}
