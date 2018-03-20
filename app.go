package main

type App struct {
	Router *mux.Router
	DB     *sql.DB
}
