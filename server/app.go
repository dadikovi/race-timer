package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/dadikovi/race-timer/server/core"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
	race   core.Race
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.ensureTableExists()

	a.race, err = core.GetRaceInstance(a.DB)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.Router.Use(prometheusMiddleware)
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
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
	race_time BIGINT,
	CONSTRAINT participant_group FOREIGN KEY(group_id) REFERENCES groups(id)
);

CREATE TABLE IF NOT EXISTS races
(
	active_group_id INTEGER,
	CONSTRAINT active_group FOREIGN KEY(active_group_id) REFERENCES groups(id)
);
`
