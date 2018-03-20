package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
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
	a.Router.HandleFunc("/api/products", a.getProducts).Methods("GET")
	a.Router.HandleFunc("/api/product", a.createProduct).Methods("POST")
	a.Router.HandleFunc("/api/product/{id:[0-9]+}", a.getProduct).Methods("GET")
	a.Router.HandleFunc("/api/product/{id:[0-9]+}", a.updateProduct).Methods("PUT")
	a.Router.HandleFunc("/api/product/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")
}

