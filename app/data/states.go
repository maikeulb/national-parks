package data

import (
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/maikeulb/national-parks/app/models"
)

func GetStates(db *sql.DB, start, count int) ([]models.State, error) {
	rows, err := db.Query(
        `SELECT id, name 
         FROM states 
         LIMIT $1 
         OFFSET $2`,
		 count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	states := []models.State{}

	for rows.Next() {
		var s models.State
		if err := rows.Scan(&s.ID, &s.Name); err != nil {
			return nil, err
		}
		states = append(states, s)
	}

	return states, nil
}

func GetState(db *sql.DB, s models.State) error {
	return db.QueryRow(
		`SELECT name 
		FROM states 
		WHERE id=$1`,
		s.ID).Scan(&s.Name)
}

func CreateState(db *sql.DB, s models.State) error {
	err := db.QueryRow(
		`INSERT INTO states(name) 
		VALUES($1) 
		RETURNING id`,
		s.Name).Scan(&s.ID)

	if err != nil {
		return err
	}

	return nil
}

func UpdateState(db *sql.DB, s models.State) error {
	_, err := db.Exec(
		`UPDATE states 
		SET name=$1 
		WHERE id=$2`,
		s.Name, s.ID)

	return err
}

func DeleteState(db *sql.DB, s models.State) error {
	_, err := db.Exec(
		`DELETE 
		FROM states 
		WHERE id=$1`,
		s.ID)

	return err
}
