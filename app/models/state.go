package models

import (
	"encoding/json"
	// "fmt"
	// "time"
)

type State struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (s State) MarshalJSON() ([]byte, error) {
	return json.Marshal(StateResponse(s))
}

func (s *State) UnmarshalJSON(data []byte) error {
	var js StateRequest
	if err := json.Unmarshal(data, &js); err != nil {
		return err
	}
	if err := js.validate(); err != nil {
		return err
	}
	*s = js.State(*s)
	return nil
}
