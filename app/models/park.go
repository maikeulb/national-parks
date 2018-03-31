package models

import (
	// "encoding/json"
	// "fmt"
	"time"
)

// type Park struct {
// 	ID          int
// 	Name        string
// 	Description string
// 	NearestCity string
// 	Visitors    int
// 	Established time.Time
// 	StateID     int
// }

type Park struct {
	ID          int        `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	NearestCity string     `json:"nearestCity,omitempty"`
	Visitors    int        `json:"visitors,omitempty"`
	Established *time.Time `json:"established,omitempty"`
	StateID     int        `json:"stateId,omitempty"`
}

// func (p Park) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(JSONPark(p))
// }

// func (p *Park) UnmarshalJSON(data []byte) error {
// 	var jp JSONPark
// 	if err := json.Unmarshal(data, &jp); err != nil {
// 		return err
// 	}
// 	// *p = jp
// 	return nil
// }
