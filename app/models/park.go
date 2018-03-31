package models

import (
	"encoding/json"
	// "fmt"
	"time"
)

type Park struct {
	ID          int
	Name        string
	Description string
	NearestCity string
	Visitors    int
	Established *time.Time
	StateID     int
}

func (p Park) MarshalJSON() ([]byte, error) {
	return json.Marshal(ParksResponse(p))
}

func (p *Park) UnmarshalJSON(data []byte) error {
	var jp ParkRequest
	if err := json.Unmarshal(data, &jp); err != nil {
		return err
	}
	if err := jp.validate(); err != nil {
		return err
	}
	*p = jp.Park(*p)
	return nil
}
