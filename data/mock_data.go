package data

import (
	"github.com/stretchr/testify/mock"
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
