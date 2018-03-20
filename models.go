package main

import "time"

type state struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type nationalPark struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	DateDesignated time.Time `json:"dateDesignated"`
	StateID        int       `json:"stateId"`
}

func (s *state) getStates(db *sql.DB, start, count int) ([]state, error) {
	return errors.New("Not implemented")
}

func (s *state) getState(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (s *state) createState(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (s *state) updateState(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (s *state) deleteState(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *nationalPark) getNationalParks(db *sql.DB, start, count int) ([]state, error) {
	return errors.New("Not implemented")
}

func (p *park) getNationalPark(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *park) createNationalPark(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *park) updateNationalPark(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *park) deleteNationalPark(db *sql.DB) error {
	return errors.New("Not implemented")
}
