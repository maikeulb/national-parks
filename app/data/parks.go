package data

import (
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/maikeulb/national-parks/app/models"
)

func GetParks(db *sql.DB, start, count, sid int) ([]models.Park, error) {
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

	parks := []models.Park{}

	for rows.Next() {
		var p models.Park
		if err := rows.Scan(
			&p.ID,
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

func GetPark(db *sql.DB, p *models.Park) error {
	return db.QueryRow(
		`SELECT id,
                name,
                description,
                nearest_city,
                visitors,
                established,
                state_id
		 FROM parks
		 WHERE id=$1 AND state_id = $2`,
		p.ID, p.StateID).Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.NearestCity,
		&p.Visitors,
		&p.Established,
		&p.StateID)
}

func CreatePark(db *sql.DB, p models.Park) error {
	err := db.QueryRow(
		`INSERT INTO parks(name, description, nearest_city, visitors, established, state_id)
						VALUES($1, $2, $3, $4, $5, $6)
						RETURNING id`,
		p.Name, p.Description, p.NearestCity, p.Visitors, p.Established, p.StateID).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}

func UpdatePark(db *sql.DB, p models.Park) error {
	_, err := db.Exec(
		`UPDATE parks
							SET name=$1
							decription=$2
							nearest_city=$3
							visitors=$4
							established=$5
							state_id=$6
							WHERE id=$7`,
		p.Name, p.Description, p.NearestCity, p.Visitors, p.Established, p.StateID, p.ID)

	return err
}

func DeletePark(db *sql.DB, p models.Park) error {
	_, err := db.Exec(
		`DELETE
								FROM parks
								WHERE id=$1 AND state_id=$2`,
		p.ID, p.StateID)

	return err
}
