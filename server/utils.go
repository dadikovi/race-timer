package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func parseRequestBody(w http.ResponseWriter, r *http.Request, dto interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return err
	}

	if err := json.Unmarshal(body, dto); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return err
	}

	return nil
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	message = strings.ReplaceAll(message, "\"", "\\\"")
	respondWithJSON(w, code, []byte(`{"error": "`+message+`"}`))
}

func respondWithJSON(w http.ResponseWriter, code int, payload []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
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
