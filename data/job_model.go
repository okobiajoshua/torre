package data

import "time"

// Skill Struct
type Skill struct {
	Name       string  `json:"name"`
	Experience string  `json:"experience"`
	Weight     float64 `json:"weight"`
}

// CompensationData struct
type CompensationData struct {
	Code        string  `json:"code"`
	Currency    string  `json:"currency"`
	MinAmount   float64 `json:"minAmount"`
	MaxAmount   float64 `json:"maxAmount"`
	Periodicity string  `json:"periodicity"`
}

// Compensation struct
type Compensation struct {
	Data    CompensationData `json:"data"`
	Visible bool             `json:"visible"`
}

// Job struct
type Job struct {
	ID            string         `json:"id"`
	Objective     string         `json:"objective"`
	Type          string         `json:"type"`
	Organizations []Organization `json:"organizations"`
	Locations     []string       `json:"locations"`
	Remote        bool           `json:"remote"`
	External      bool           `json:"external"`
	Deadline      time.Time      `json:"deadline"`
	Status        string         `json:"status"`
	Compensation  Compensation   `json:"compensation"`
	Skills        []Skill        `json:"skills"`
	Description   string         `json:"serpTags.description"`
}
