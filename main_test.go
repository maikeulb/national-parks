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
	if body := response.Body.String(); body != "{'data':[]}" {
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

func TestGetNonExistentPark(t *testing.T) {
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

func TestCreatePark(t *testing.T) {
	clearTable()
	addStates(1)

	payload := []byte(`{"name":"test park"}`)

	req, _ := http.NewRequest("POST", "/api/states/1/park", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "test park" {
		t.Errorf("Expected state name to be 'test state'. Got '%v'", m["name"])
	}
}

// func TestGetPark(t *testing.T) {
// 	clearTable()
// 	addStatesAndParks(1)

// req, _ := http.NewRequest("GET", "/api/states/1/parks/1", nil)
// response := executeRequest(req)

// checkResponseCode(t, http.StatusOK, response.Code)
// }

// func TestUpdatePark(t *testing.T) {
// 	clearTable()
// 	addStates(1)
// 	addParks(1)

// 	req, _ := http.NewRequest("GET", "/api/states/1/parks/1", nil)
// 	response := executeRequest(req)

// 	var originalPark map[string]interface{}
// 	json.Unmarshal(response.Body.Bytes(), &originalPark)

// 	payload := []byte(
// 		`{"name":"test park - updated name"}`)

// 	req, _ = http.NewRequest("PATCH", "/api/states/1/parks/1", bytes.NewBuffer(payload))
// 	response = executeRequest(req)

// 	checkResponseCode(t, http.StatusOK, response.Code)

// 	var m map[string]interface{}
// 	json.Unmarshal(response.Body.Bytes(), &m)

// 	if m["name"] == originalPark["name"] {
// 		t.Errorf("Expected the name to change from '%v' to '%v'. Got '%v'", originalPark["name"], m["name"], m["name"])
// 	}
// }

// func TestDeletePark(t *testing.T) {
// 	clearTable()
// 	addStates(1)
// 	addParks(1)

// 	req, _ := http.NewRequest("GET", "/api/states/1/parks/1", nil)
// 	response := executeRequest(req)
// 	checkResponseCode(t, http.StatusOK, response.Code)

// 	req, _ = http.NewRequest("DELETE", "/api/states/1/parks/1", nil)
// 	response = executeRequest(req)

// 	checkResponseCode(t, http.StatusOK, response.Code)

// 	req, _ = http.NewRequest("GET", "/api/states/1/parks/1", nil)
// 	response = executeRequest(req)
// 	checkResponseCode(t, http.StatusNotFound, response.Code)
// }

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
	a.DB.Exec("DELETE FROM parks")
	a.DB.Exec("ALTER SEQUENCE states_id_seq RESTART WITH 1")
	a.DB.Exec("ALTER SEQUENCE parks_id_seq RESTART WITH 1")
}

func addStates(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec(`
        INSERT INTO states(name) 
            VALUES($1)`,
			"State "+strconv.Itoa(i))
	}
}

func addStatesAndParks(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec(`
        INSERT INTO states(name) 
            VALUES($1)`,
			"State "+strconv.Itoa(i))
		a.DB.Exec(`
        INSERT INTO parks(name, description, nearest_city, visitors, state_id) 
            VALUES($1,$2,$3,$4,$5)`,
			"State "+strconv.Itoa(i),
			"Description "+strconv.Itoa(i),
			"Nearest_city "+strconv.Itoa(i),
			i,
			i)
	}
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS states
        (
            id SERIAL,
            name varchar(50) NOT NULL,
            CONSTRAINT products_pkey PRIMARY KEY (id)
        );
        
        CREATE TABLE IF NOT EXISTS parks 
        (
            id serial,
            name varchar(50) NOT NULL,
            description TEXT NOT NULL,
            nearest_city varchar(50) NOT NULL,
            visitors integer NOT NULL,
            established timestamp NULL,
            state_id integer NULL,
            CONSTRAINT parks_pkey PRIMARY KEY (id),
            FOREIGN KEY (state_id) REFERENCES states(id)
        );

        `
