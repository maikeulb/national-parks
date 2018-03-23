package models

import (
	"time"
)

type Park struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	NearestCity string    `json:"nearestCity"`
	Visitors    int       `json:"visitors"`
	Established time.Time `json:"established"`
	StateID     int       `json:"stateId"`
}
