package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

type RAWROW map[string]interface{}

func (a *App) Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/segments", a.createSegmentMiddleware).Methods("POST")
	a.Router.HandleFunc("/segments", a.fetchAllSegmentMiddleware).Methods("GET")
	a.Router.HandleFunc("/groups", a.createGroupMiddleware).Methods("POST")
}

func (a *App) createSegmentMiddleware(w http.ResponseWriter, r *http.Request) {
	var segmentEndpoint SegmentEndpoint
	segmentEndpoint.create(w, r, a.DB)
}

func (a *App) fetchAllSegmentMiddleware(w http.ResponseWriter, r *http.Request) {
	var segmentEndpoint SegmentEndpoint
	segmentEndpoint.fetchAll(w, r, a.DB)
}

func (a *App) createGroupMiddleware(w http.ResponseWriter, r *http.Request) {
	var groupEndpoint GroupEndpoint
	groupEndpoint.create(w, r, a.DB)
}
