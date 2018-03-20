package main

import "time"

type State struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Park struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	DateDesignated time.Tim `json:"dateDesignated"`
	StateID        int      `json:"stateId"`
}
