package main

import (
	"database/sql"
)

func getStates(db *sql.DB, start, count int) ([]state, error) {
	rows, err := db.Query(
		"SELECT id, name FROM states LIMIT $1 OFFSET $2",
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

func (s *state) getState(db *sql.DB) error {
	return db.QueryRow("SELECT name FROM states WHERE id=$1",
		s.ID).Scan(&s.Name)
}

func (s *state) createState(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO states(name) VALUES($1) RETURNING id",
		s.Name).Scan(&s.ID)

	if err != nil {
		return err
	}

	return nil
}

func (s *state) updateState(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE states SET name=$1 WHERE id=$2",
			s.Name, s.ID)

	return err
}

func (s *state) deleteState(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM states WHERE id=$1", s.ID)

	return err
}

// func getNationalParks(db *sql.DB, start, count int) ([]state, error) {
// 	return nil, errors.New("Not implemented")
// }

// func (p *nationalPark) getNationalPark(db *sql.DB) error {
// 	return errors.New("Not implemented")
// }

// func (p *nationalPark) createNationalPark(db *sql.DB) error {
// 	return errors.New("Not implemented")
// }

// func (p *nationalPark) updateNationalPark(db *sql.DB) error {
// 	return errors.New("Not implemented")
// }

// func (p *nationalPark) deleteNationalPark(db *sql.DB) error {
// 	return errors.New("Not implemented")
// }
