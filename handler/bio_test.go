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

func TestBioDataFetchWithoutUsernameShouldReturnBadRequest(t *testing.T) {
	m := &data.MockData{}
	b := NewBioHandler(m)

	req := httptest.NewRequest(http.MethodGet, "/torre/bio/", nil)
	res := httptest.NewRecorder()

	b.GetBioByUsername(res, req)
	assert.Equal(t, http.StatusBadRequest, res.Code, "Should return a 400 bad request")
}

func TestBioDataFetchWithWrongUsernameShouldReturnUnprocessableEntity(t *testing.T) {
	m := &data.MockData{}
	b := NewBioHandler(m)
	r := mux.NewRouter()
	r.HandleFunc("/torre/bio/{username}", b.GetBioByUsername).Methods(http.MethodGet)

	req := httptest.NewRequest(http.MethodGet, "/torre/bio/xyz", nil)
	res := httptest.NewRecorder()

	m.On("GetBioDataByUsername", mock.Anything).Return(nil, fmt.Errorf("not found"))

	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusUnprocessableEntity, res.Code, "Should return a 422 unprocessable entity")
}

func TestBioDataFetchWithUsernameShouldReturnSuccess(t *testing.T) {
	m := &data.MockData{}
	b := NewBioHandler(m)
	r := mux.NewRouter()
	r.HandleFunc("/torre/bio/{username}", b.GetBioByUsername).Methods(http.MethodGet)

	req := httptest.NewRequest(http.MethodGet, "/torre/bio/xyz", nil)
	res := httptest.NewRecorder()

	m.On("GetBioDataByUsername", mock.Anything).Return(&data.BioData{}, nil)

	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code, "Should return a 200 status OK")
}
