package main

import "time"

// Airline holds info for single record in fleet table
type Airline struct {
	ID          string    `json:"id"`
	FullName    string    `json:"full_name"`
	IsActivated bool      `json:"is_activated"`
	LastUpdated time.Time `json:"last_updated"`
}

// Activate holds the instructions on whether to activate an airline or not
type Activate struct {
	Activate bool `json:"activate"`
}

//Airlines holds a map of Airlines indexed by airline ID
type Airlines map[string]Airline
