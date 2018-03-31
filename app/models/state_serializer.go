package models

import (
	"errors"
	// "time"
)

type StateRequest struct {
	Name string `json:"name,omitempty"`
}

func (js StateRequest) State(s State) State {
	s.Name = js.Name

	return s
}

func (js *StateRequest) validate() error {
	if js.Name == "" {
		return errors.New("Name is required")
	}

	return nil
}

type StateResponse struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func StatesResponse(s State) StateResponse {
	var js StateResponse
	js.ID = s.ID
	js.Name = s.Name

	return js
}
