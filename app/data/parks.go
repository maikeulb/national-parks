package data

import (
	"database/sql"
	"fmt"
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
	fmt.Println(p.Established)
	_, err := db.Exec(
		`UPDATE parks as p
         SET name=case 
                  when $1='' then p.name 
                  else $1
            end,
            description=case
                when $2='' then p.description
                else $2
            end,
            nearest_city=case
                when $3='' then p.nearest_city
                else $3
            end,
            visitors=case
                when $4=0 then p.visitors
                else $4
            end,
            established=p.established
		WHERE id=$5`,
		p.Name, p.Description, p.NearestCity, p.Visitors, p.ID)

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
