package data

import (
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/maikeulb/national-parks/app/models"
)

func GetParks(db *sql.DB, start, count, sid int) ([]models.Park, error) {
	rows, err := db.Query(
		`SELECT * FROM parks
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

// func (p *models.Park) createState(db *sql.DB) error {
//	err := db.QueryRow(
//		`INSERT INTO parks(name, description, nearest_city, visitors, established, state_id)
//					VALUES($1, $2, $3, $4, $5, $6) RETURNING id`,
//		p.Name, p.Description, p.NearestCity, p.Visitors, p.Established, p.StateID).Scan(&p.ID)

//	if err != nil {
//		return err
//	}

//	return nil
// }

// func (p *Park) updatePark(db *sql.DB) error {
//	return errors.New("Not implemented")
// }

// func (p *Park) deletePark(db *sql.DB) error {
//	return errors.New("Not implemented")
// }
