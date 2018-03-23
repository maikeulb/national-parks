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

func getParks(db *sql.DB, start, count, sid int) ([]park, error) {
	rows, err := db.Query(
		`SELECT *
				FROM parks
				WHERE state_id = $1
				LIMIT $2
				OFFSET $3`,
		sid, count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	parks := []park{}

	for rows.Next() {
		var p park
		if err := rows.Scan(&p.ID,
			&p.Name,
			&p.Description,
			&p.NearestCity,
			&p.Visitors,
			&p.Established,
			&p.StateID); err != nil {
			return nil, err
		}
		parks = append(parks, p)
	}

	return parks, nil
}

// func (p *Park) getPark(db *sql.DB) error {
//	return errors.New("Not implemented")
// }

func (p *park) createState(db *sql.DB) error {
	err := db.QueryRow(
		`INSERT INTO parks(name, description, nearest_city, visitors, established, state_id)
					VALUES($1, $2, $3, $4, $5, $6) RETURNING id`,
		p.Name, p.Description, p.NearestCity, p.Visitors, p.Established, p.StateID).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}

// func (p *Park) updatePark(db *sql.DB) error {
//	return errors.New("Not implemented")
// }

// func (p *Park) deletePark(db *sql.DB) error {
//	return errors.New("Not implemented")
// }
