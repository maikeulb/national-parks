package main_test

import (
	"bytes"
	"encoding/json"
	"github.com/maikeulb/national-parks/app"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

var a app.App

func TestMain(m *testing.M) {
	a = app.App{}

	a.Initialize(
		os.Getenv("TEST_DB_HOST"),
		os.Getenv("TEST_DB_PORT"),
		os.Getenv("TEST_DB_USERNAME"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_DB"))
	ensureTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/api/states", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetNonExistentState(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/api/states/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "State not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Invalid state ID'. Got '%s'", m["error"])
	}
}

func TestCreateState(t *testing.T) {
	clearTable()

	payload := []byte(`{"name":"test state"}`)

	req, _ := http.NewRequest("POST", "/api/states", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "test state" {
		t.Errorf("Expected state name to be 'test state'. Got '%v'", m["name"])
	}

}

func TestGetState(t *testing.T) {
	clearTable()
	addStates(1)

	req, _ := http.NewRequest("GET", "/api/states/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdateState(t *testing.T) {
	clearTable()
	addStates(1)

	req, _ := http.NewRequest("GET", "/api/states/1", nil)
	response := executeRequest(req)
	var originalState map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalState)

	payload := []byte(`{"name":"test state - updated name"}`)

	req, _ = http.NewRequest("PUT", "/api/states/1", bytes.NewBuffer(payload))
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] == originalState["name"] {
		t.Errorf("Expected the name to change from '%v' to '%v'. Got '%v'", originalState["name"], m["name"], m["name"])
	}
}

func TestDeleteState(t *testing.T) {
	clearTable()
	addStates(1)

	req, _ := http.NewRequest("GET", "/api/states/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/api/states/1", nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/api/states/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM states")
	a.DB.Exec("ALTER SEQUENCE states_id_seq RESTART WITH 1")
}

func addStates(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO states(name) VALUES($1)", "State "+strconv.Itoa(i))
	}
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS states
        (
            id SERIAL,
            name TEXT NOT NULL,
            CONSTRAINT states_pkey PRIMARY KEY (id)
        )`
