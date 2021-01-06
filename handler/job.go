package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/torre/data"
)

// Job struct handler
type Job struct {
	d data.Data
}

// NewJobHandler returns a job handler
func NewJobHandler(d data.Data) *Job {
	return &Job{d: d}
}

// GetJobByID method
func (j *Job) GetJobByID(w http.ResponseWriter, r *http.Request) {
	id := strings.Trim(mux.Vars(r)["id"], " ")
	if id == "" {
		http.Error(w, "no username found", http.StatusBadRequest)
		return
	}
	job, err := j.d.GetJobByID(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "error occured", http.StatusUnprocessableEntity)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(job)
}
