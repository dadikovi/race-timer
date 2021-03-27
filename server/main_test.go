package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dadikovi/race-timer/server/core"
)

var a App

func TestMain(m *testing.M) {
	a.Initialize(
		os.Getenv("RACE_TIMER_DB_USER"),
		os.Getenv("RACE_TIMER_DB_PASSWORD"),
		os.Getenv("RACE_TIMER_DB_NAME"))

	code := m.Run()
	clearTable()
	os.Exit(code)
}

func refreshRace() {
	var err error
	a.race, err = core.GetRaceInstance(a.DB)

	if err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	if _, err := a.DB.Exec("DELETE FROM participants;"); err != nil {
		log.Fatal(err)
	}
	if _, err := a.DB.Exec("DELETE FROM races; INSERT INTO races (active_group_id) VALUES (NULL);"); err != nil {
		log.Fatal(err)
	}
	if _, err := a.DB.Exec("DELETE FROM groups"); err != nil {
		log.Fatal(err)
	}
	if _, err := a.DB.Exec("ALTER SEQUENCE groups_id_seq RESTART WITH 1"); err != nil {
		log.Fatal(err)
	}
	if _, err := a.DB.Exec("DELETE FROM segments"); err != nil {
		log.Fatal(err)
	}
	if _, err := a.DB.Exec("ALTER SEQUENCE segments_id_seq RESTART WITH 1"); err != nil {
		log.Fatal(err)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}
