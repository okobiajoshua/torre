package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/torre/data"
)

func TestJobFetchWithoutIDShouldReturnBadRequest(t *testing.T) {
	m := &data.MockData{}
	j := NewJobHandler(m)

	req := httptest.NewRequest(http.MethodGet, "/torre/job/", nil)
	res := httptest.NewRecorder()

	j.GetJobByID(res, req)
	assert.Equal(t, http.StatusBadRequest, res.Code, "Should return a 400 bad request")
}

func TestJobFetchWithWrongIDShouldReturnUnprocessableEntity(t *testing.T) {
	m := &data.MockData{}
	j := NewJobHandler(m)

	r := mux.NewRouter()
	r.HandleFunc("/torre/job/{id}", j.GetJobByID).Methods(http.MethodGet)

	req := httptest.NewRequest(http.MethodGet, "/torre/job/xyz", nil)
	res := httptest.NewRecorder()

	m.On("GetJobByID", mock.Anything).Return(nil, fmt.Errorf("not found"))

	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusUnprocessableEntity, res.Code, "Should return a 422 unprocessable entity")
}

func TestJobFetchWithIDShouldReturnSuccess(t *testing.T) {
	m := &data.MockData{}
	j := NewJobHandler(m)

	r := mux.NewRouter()
	r.HandleFunc("/torre/job/{id}", j.GetJobByID).Methods(http.MethodGet)

	req := httptest.NewRequest(http.MethodGet, "/torre/job/xyz", nil)
	res := httptest.NewRecorder()

	m.On("GetJobByID", mock.Anything).Return(&data.Job{}, nil)

	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code, "Should return a 200 status OK")
}
