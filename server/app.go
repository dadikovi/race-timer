package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) getTimeTable(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Time table")
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/results", a.getTimeTable).Methods("GET")
}
