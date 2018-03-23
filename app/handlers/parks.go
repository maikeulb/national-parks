package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/maikeulb/national-parks/app/data"
	"github.com/maikeulb/national-parks/app/models"
)

func GetParks(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))
	vars := mux.Vars(r)

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	sid, err := strconv.Atoi(vars["sid"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid state ID")
		return
	}

	parks, err := data.GetParks(db, start, count, sid)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, parks)
}

func GetPark(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid park ID")
		return
	}

	sid, err := strconv.Atoi(vars["sid"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid state ID")
		return
	}

	p := models.Park{ID: id, StateID: sid}
	if err := data.GetPark(db, p); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Park not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func CreatePark(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sid, err := strconv.Atoi(vars["sid"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid state ID")
		return
	}
	p := models.Park{StateID: sid}

	// var p models.Park
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := data.CreatePark(db, p); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}

func UpdatePark(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid park ID")
		return
	}

	sid, err := strconv.Atoi(vars["sid"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid state ID")
		return
	}

	var p models.Park
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	p.ID = id
	p.StateID = sid

	if err := data.UpdatePark(db, p); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func DeletePark(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid park ID")
		return
	}

	p := models.Park{ID: id}
	if err := data.DeletePark(db, p); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
