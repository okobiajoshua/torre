package data

import (
	"encoding/json"
	"net/http"
)

// Fetch fetches data from torre APIs
type Fetch struct {
}

// NewFetch returns a fetch struct
func NewFetch() *Fetch {
	return &Fetch{}
}

// GetBioDataByUsername method
func (f *Fetch) GetBioDataByUsername(username string) (*BioData, error) {
	res, err := http.Get("https://bio.torre.co/api/bios/" + username)
	if err != nil {
		return nil, err
	}
	var bd BioData
	err = json.NewDecoder(res.Body).Decode(&bd)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	return &bd, nil
}
