package models

import (
	// "errors"
	"time"
)

type ParkRequest struct {
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	NearestCity string     `json:"nearestCity,omitempty"`
	Visitors    int        `json:"visitors,omitempty"`
	Established *time.Time `json:"established,omitempty"`
}

func (jp ParkRequest) Park(p Park) Park {
	p.Name = jp.Name
	p.Description = jp.Description
	p.NearestCity = jp.NearestCity
	p.Visitors = jp.Visitors
	p.Established = jp.Established

	return p
}

func (jp *ParkRequest) validate() error {
	// if jp.FolloweeID <= 0 {
	// return errors.New("FolloweeID must not be empty")
	// }

	return nil
}

type ParkResponse struct {
	ID          int        `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	NearestCity string     `json:"nearestCity,omitempty"`
	Visitors    int        `json:"visitors,omitempty"`
	Established *time.Time `json:"established,omitempty"`
	StateID     int        `json:"stateId,omitempty"`
}

func ParksResponse(p Park) ParkResponse {
	var jp ParkResponse
	jp.ID = p.ID
	jp.StateID = p.StateID
	jp.Name = p.Name
	jp.Description = p.Description
	jp.NearestCity = p.NearestCity
	jp.Visitors = p.Visitors

	return jp
}
