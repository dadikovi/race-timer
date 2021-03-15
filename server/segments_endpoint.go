package main

import (
	"database/sql"
	"io/ioutil"
	"net/http"

	"github.com/dadikovi/race-timer/server/core"
)

type SegmentEndpoint struct{}

func (se *SegmentEndpoint) create(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var body, bodyReadError = ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if bodyReadError != nil {
		respondWithError(w, http.StatusBadRequest, bodyReadError.Error())
	}

	var s, err = core.MakeSegment(string(body))

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	if err := s.Save(db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var json, _ = s.ToJson()
	respondWithJSON(w, http.StatusCreated, json)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, []byte(`{"error": `+message+`}`))
}

func respondWithJSON(w http.ResponseWriter, code int, payload []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(payload)
}
