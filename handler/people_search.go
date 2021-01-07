package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/torre/data"
	"github.com/torre/dto"
	"github.com/torre/model"
)

// JobSearch struct handler
type PeopleSearch struct {
	d data.Data
}

// NewJobSearchHandler returns a JobSearch struct
func NewPeopleSearchHandler(d data.Data) *PeopleSearch {
	return &PeopleSearch{d: d}
}

// Find handler returns a list of job that fulfils the search criteria
func (ps *PeopleSearch) Find(w http.ResponseWriter, r *http.Request) {
	var sp dto.SearchParam
	var page model.Page
	json.NewDecoder(r.Body).Decode(&sp)
	defer r.Body.Close()

	if size := mux.Vars(r)["size"]; size != "" {
		page.Size, _ = strconv.Atoi(size)
	}
	if offset := mux.Vars(r)["offset"]; offset != "" {
		page.Offset, _ = strconv.Atoi(offset)
	}
	aggregate := mux.Vars(r)["offset"]
	page.Aggregate, _ = strconv.ParseBool(aggregate)
	if page.Size <= 0 {
		page.Size = 10
	}

	jobs, err := ps.d.GetPeople(sp, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}
