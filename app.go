package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(host string, port int, user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":5000", a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/api/states", a.getStates).Methods("GET")
	a.Router.HandleFunc("/api/states", a.createState).Methods("POST")
	a.Router.HandleFunc("/api/states/{id:[0-9]+}", a.getState).Methods("GET")
	a.Router.HandleFunc("/api/states/{id:[0-9]+}", a.updateState).Methods("PUT")
	a.Router.HandleFunc("/api/states/{id:[0-9]+}", a.deleteState).Methods("DELETE")
	a.Router.HandleFunc("/api/states/{sid}/parks", a.getParks).Methods("GET")
	// a.Router.HandleFunc("/api/states/{sid}/national_parks", a.createNationalPark).Methods("POST")
	// a.Router.HandleFunc("/api/states/{sid}/national_parks/{id:[0-9]+}", a.getNationalPark).Methods("GET")
	// a.Router.HandleFunc("/api/states/{sid}/national_parks/{id:[0-9]+}", a.updateNationalPark).Methods("PUT")
	// a.Router.HandleFunc("/api/states/{sid}/national_parks/{id:[0-9]+}", a.deleteNationalPark).Methods("DELETE")
}
