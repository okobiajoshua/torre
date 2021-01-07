package data

import (
	"github.com/torre/dto"
	"github.com/torre/model"
)

// Data interface
type Data interface {
	GetBioDataByUsername(username string) (*BioData, error)
	GetJobByID(jobID string) (*Job, error)
	GetJobs(jsp dto.SearchParam, page model.Page) ([]Job, error)
	GetPeople(jsp dto.SearchParam, page model.Page) ([]Person, error)
}
