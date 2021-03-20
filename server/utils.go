package main

import "net/http"

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, []byte(`{"error": `+message+`}`))
}

func respondWithJSON(w http.ResponseWriter, code int, payload []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(payload)
}

type RAWROW map[string]interface{}
