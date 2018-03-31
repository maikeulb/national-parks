package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maikeulb/national-parks/app/handlers"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(host, port, user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
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
	a.Router.Use(limitMiddleware)
}
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/api/states", a.GetStates).Methods("GET")
	a.Router.HandleFunc("/api/states", a.CreateState).Methods("POST")
	a.Router.HandleFunc("/api/states/{id:[0-9]+}", a.GetState).Methods("GET")
	a.Router.HandleFunc("/api/states/{id:[0-9]+}", a.UpdateState).Methods("PUT")
	a.Router.HandleFunc("/api/states/{id:[0-9]+}", a.DeleteState).Methods("DELETE")
	a.Router.HandleFunc("/api/states/{sid}/parks", a.GetParks).Methods("GET")
	a.Router.HandleFunc("/api/states/{sid}/parks", a.CreatePark).Methods("POST")
	a.Router.HandleFunc("/api/states/{sid}/parks/{id:[0-9]+}", a.GetPark).Methods("GET")
	a.Router.HandleFunc("/api/states/{sid}/parks/{id:[0-9]+}", a.UpdatePark).Methods("PATCH")
	a.Router.HandleFunc("/api/states/{sid}/parks/{id:[0-9]+}", a.DeletePark).Methods("DELETE")
	a.Router.Use(limitMiddleware)
}

func (a *App) GetStates(w http.ResponseWriter, r *http.Request) {
	handlers.GetStates(a.DB, w, r)
}

func (a *App) CreateState(w http.ResponseWriter, r *http.Request) {
	handlers.CreateState(a.DB, w, r)
}

func (a *App) GetState(w http.ResponseWriter, r *http.Request) {
	handlers.GetState(a.DB, w, r)
}

func (a *App) UpdateState(w http.ResponseWriter, r *http.Request) {
	handlers.UpdateState(a.DB, w, r)
}

func (a *App) DeleteState(w http.ResponseWriter, r *http.Request) {
	handlers.DeleteState(a.DB, w, r)
}

func (a *App) GetParks(w http.ResponseWriter, r *http.Request) {
	handlers.GetParks(a.DB, w, r)
}

func (a *App) CreatePark(w http.ResponseWriter, r *http.Request) {
	handlers.CreatePark(a.DB, w, r)
}

func (a *App) GetPark(w http.ResponseWriter, r *http.Request) {
	handlers.GetPark(a.DB, w, r)
}

func (a *App) UpdatePark(w http.ResponseWriter, r *http.Request) {
	handlers.UpdatePark(a.DB, w, r)
}

func (a *App) DeletePark(w http.ResponseWriter, r *http.Request) {
	handlers.DeletePark(a.DB, w, r)
}
