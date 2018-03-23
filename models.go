package main

import (
	"time"
)

type state struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type park struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	DateDesignated time.Time `json:"dateDesignated"`
	StateID        int       `json:"stateId"`
}
