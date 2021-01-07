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

func TestJobSearchWithoutParamsShouldReturnSuccess(t *testing.T) {
	m := &data.MockData{}
	jsh := NewJobSearchHandler(m)

	req := httptest.NewRequest(http.MethodPost, "/torre/search/job/", nil)
	res := httptest.NewRecorder()
	m.On("GetJobs", mock.Anything, mock.Anything).Return([]data.Job{}, nil)

	jsh.Find(res, req)
	assert.Equal(t, http.StatusOK, res.Code, "Should return a 200 status OK")
}

func TestJobSearchWithInternalErrorShouldReturnInternalServerError(t *testing.T) {
	m := &data.MockData{}
	jsh := NewJobSearchHandler(m)

	req := httptest.NewRequest(http.MethodGet, "/torre/search/job", nil)
	res := httptest.NewRecorder()

	m.On("GetJobs", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("not found"))

	jsh.Find(res, req)
	assert.Equal(t, http.StatusInternalServerError, res.Code, "Should return a 500 internal server error")
}

func TestJobFetchWithParamsShouldReturnSuccess(t *testing.T) {
	m := &data.MockData{}
	jsh := NewJobSearchHandler(m)

	r := mux.NewRouter()
	r.HandleFunc("/torre/search/job/", jsh.Find).Methods(http.MethodPost)

	req := httptest.NewRequest(http.MethodPost, "/torre/search/job/?size=10&offset=0", nil)
	res := httptest.NewRecorder()

	m.On("GetJobs", mock.Anything, mock.Anything).Return([]data.Job{}, nil)

	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code, "Should return a 200 status OK")
}
