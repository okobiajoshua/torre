package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/torre/data"
)

// Bio handler
type Bio struct {
	d data.Data
}

// NewBioHandler method
func NewBioHandler(d data.Data) *Bio {
	return &Bio{d: d}
}

// GetBioByUsername method
func (b *Bio) GetBioByUsername(w http.ResponseWriter, r *http.Request) {
	uname := strings.Trim(mux.Vars(r)["username"], " ")
	if uname == "" {
		http.Error(w, "no username found", http.StatusBadRequest)
		return
	}
	bd, err := b.d.GetBioDataByUsername(uname)
	if err != nil {
		log.Println(err)
		http.Error(w, "error occured", http.StatusUnprocessableEntity)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(bd)
}
