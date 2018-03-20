package main

import "time"

type state struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type nationalPark struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	DateDesignated time.Tim `json:"dateDesignated"`
	StateID        int      `json:"stateId"`
}

func getStates(db *sql.DB, start, count int) ([]state, error) {
	rows, err := db.Query(
		"",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	states := []state{}

	for rows.Next() {
		var s state
		if err := rows.Scan(&s.ID, &s.Name); err != nil {
			return nil, err
		}
		states = append(states, s)
	}

	return states, nil
}

func (s *State) getState(db *sql.DB) error {
	return db.QueryRow(
		"",
		s.ID).Scan(&s.Name)
}

func (s *state) createState(db *sql.DB) error {
	err := db.QueryRow(
		"",
		s.Name).Scan(&s.ID)

	if err != nil {
		return err
	}

	return nil
}

func (s *state) updateState(db *sql.DB) error {
	_, err :=
		db.Exec(
			"",
			s.Name, s.ID)

	return err
}

func (s *state) deleteState(db *sql.DB) error {
	_, err := db.Exec(
		"",
		s.ID)

	return err
}

