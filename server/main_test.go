package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a.Initialize(
		os.Getenv("RACE_TIMER_DB_USER"),
		os.Getenv("RACE_TIMER_DB_PASSWORD"),
		os.Getenv("RACE_TIMER_DB_NAME"))

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	if _, err := a.DB.Exec("DELETE FROM participants;"); err != nil {
		log.Fatal(err)
	}
	if _, err := a.DB.Exec("DELETE FROM races"); err != nil {
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

const tableCreationQuery = `
CREATE TABLE IF NOT EXISTS segments
(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS groups
(
    id SERIAL PRIMARY KEY,
	start TIMESTAMP,
	segment_id INTEGER,
	CONSTRAINT group_segment FOREIGN KEY(segment_id) REFERENCES segments(id)
);

CREATE TABLE IF NOT EXISTS participants
(
    start_number INTEGER PRIMARY KEY,
	finish TIMESTAMP,
	group_id INTEGER,
	race_time INTEGER,
	CONSTRAINT participant_group FOREIGN KEY(group_id) REFERENCES groups(id)
);

CREATE TABLE IF NOT EXISTS races
(
	active_group_id INTEGER,
	CONSTRAINT active_group FOREIGN KEY(active_group_id) REFERENCES groups(id)
);
`
