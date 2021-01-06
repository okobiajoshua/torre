package handler

import (
	"net/http"

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

func (j *Job) GetJobByID(w http.ResponseWriter, r *http.Request) {

}
