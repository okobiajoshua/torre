package data

import (
	"github.com/stretchr/testify/mock"
	"github.com/torre/dto"
	"github.com/torre/model"
)

// MockData struct
type MockData struct {
	mock.Mock
}

// GetBioDataByUsername mock method
func (m *MockData) GetBioDataByUsername(username string) (*BioData, error) {
	args := m.Called(username)
	bd := args.Get(0)
	if bd != nil {
		return bd.(*BioData), args.Error(1)
	}
	return nil, args.Error(1)
}

// GetJobByID mock method
func (m *MockData) GetJobByID(jobID string) (*Job, error) {
	args := m.Called(jobID)
	job := args.Get(0)
	if job != nil {
		return job.(*Job), args.Error(1)
	}
	return nil, args.Error(1)
}

// GetJobs mock method
func (m *MockData) GetJobs(param dto.SearchParam, page model.Page) ([]Job, error) {
	args := m.Called(param, page)
	jobs := args.Get(0)
	if jobs != nil {
		return jobs.([]Job), args.Error(1)
	}
	return nil, args.Error(1)
}

// GetPeople mock method
func (m *MockData) GetPeople(param dto.SearchParam, page model.Page) ([]Person, error) {
	args := m.Called(param, page)
	p := args.Get(0)
	if p != nil {
		return p.([]Person), args.Error(1)
	}
	return nil, args.Error(1)
}
