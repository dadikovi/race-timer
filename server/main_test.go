package main

import (
	"log"
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
	a.DB.Exec("DELETE FROM segments")
	a.DB.Exec("ALTER SEQUENCE segments_id_seq RESTART WITH 1")

	a.DB.Exec("DELETE FROM groups")
	a.DB.Exec("ALTER SEQUENCE groups_id_seq RESTART WITH 1")

	a.DB.Exec("DELETE FROM participants")
	a.DB.Exec("ALTER SEQUENCE participants_id_seq RESTART WITH 1")
}

const tableCreationQuery = `
CREATE TABLE IF NOT EXISTS segments
(
    id SERIAL,
    name TEXT NOT NULL,
    CONSTRAINT segments_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS groups
(
    id SERIAL,
    name TEXT NOT NULL,
	start TIMESTAMP,
	segment_id SERIAL,
	CONSTRAINT group_segment FOREIGN KEY(segment_id) REFERENCES segments(id),
    CONSTRAINT groups_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS participants
(
    id SERIAL,
	finish TIMESTAMP,
	group_id SERIAL,
	CONSTRAINT participant_group FOREIGN KEY(group_id) REFERENCES groups(id),
    CONSTRAINT participants_key PRIMARY KEY (id)
);
`
