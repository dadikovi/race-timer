package main

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/segments", a.createSegment).Methods("POST")
	a.Router.HandleFunc("/segments", a.fetchAllSegment).Methods("GET")
	a.Router.HandleFunc("/groups", a.createGroup).Methods("POST")
	a.Router.HandleFunc("/groups/active", a.startActiveGroup).Methods("PUT")
	a.Router.HandleFunc("/participants", a.registerParticipant).Methods("POST")
	a.Router.HandleFunc("/participants/{startNumber}", a.participantFinished).Methods("PUT")
}
