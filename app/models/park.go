package models

import (
	"time"
)

type Park struct {
	ID          int       `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	NearestCity string    `json:"nearestCity,omitempty"`
	Visitors    int       `json:"visitors,omitempty"`
	Established time.Time `json:"established,omitempty"`
	StateID     int       `json:"stateId,omitempty"`
}
