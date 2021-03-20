package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func parseRequestBody(w http.ResponseWriter, r *http.Request, dto interface{}) {
	var body, bodyReadError = ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if bodyReadError != nil {
		respondWithError(w, http.StatusBadRequest, bodyReadError.Error())
	}

	if err := json.Unmarshal(body, dto); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, []byte(`{"error": `+message+`}`))
}

func respondWithJSON(w http.ResponseWriter, code int, payload []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(payload)
}

func respondWithDto(w http.ResponseWriter, dto interface{}) {
	j, err := json.Marshal(dto)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJSON(w, http.StatusOK, j)
}

func respondWithServerError(w http.ResponseWriter, err error) {
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
}

func respondWithClientError(w http.ResponseWriter, err error) {
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
}
