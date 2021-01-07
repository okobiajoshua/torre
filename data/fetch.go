package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/torre/dto"
	"github.com/torre/model"
)

// Fetch fetches data from torre APIs
type Fetch struct {
}

// NewFetch returns a fetch struct
func NewFetch() *Fetch {
	return &Fetch{}
}

// GetBioDataByUsername method
func (f *Fetch) GetBioDataByUsername(username string) (*BioData, error) {
	res, err := http.Get("https://bio.torre.co/api/bios/" + username)
	if err != nil {
		return nil, err
	}
	var bd BioData
	err = json.NewDecoder(res.Body).Decode(&bd)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	return &bd, nil
}

// GetJobByID method
func (f *Fetch) GetJobByID(jobID string) (*Job, error) {
	res, err := http.Get("https://torre.co/api/opportunities/" + jobID)
	if err != nil {
		return nil, err
	}
	var job Job
	err = json.NewDecoder(res.Body).Decode(&job)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	return &job, nil
}

// GetJobs method
func (f *Fetch) GetJobs(sp dto.SearchParam, page model.Page) ([]Job, error) {
	body, err := searchParamToByte(sp)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://search.torre.co/opportunities/_search/?offset=%d&size=%d&aggregate=%t", page.Offset, page.Size, page.Aggregate)

	var res *http.Response

	if body != nil {
		res, err = http.Post(url, "application/json", body)
	} else {
		res, err = http.Post(url, "application/json", nil)
	}
	if err != nil {
		return nil, err
	}

	var j Jobs
	err = json.NewDecoder(res.Body).Decode(&j)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	return j.Results, nil
}

// GetPeople method
func (f *Fetch) GetPeople(sp dto.SearchParam, page model.Page) ([]Person, error) {
	body, err := searchParamToByte(sp)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://search.torre.co/people/_search/?offset=%d&size=%d&aggregate=%t", page.Offset, page.Size, page.Aggregate)

	var res *http.Response

	if body != nil {
		res, err = http.Post(url, "application/json", body)
	} else {
		res, err = http.Post(url, "application/json", nil)
	}
	if err != nil {
		return nil, err
	}

	var p People
	err = json.NewDecoder(res.Body).Decode(&p)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	return p.Results, nil
}

func searchParamToByte(sp dto.SearchParam) (*bytes.Buffer, error) {
	m := map[string]interface{}{}
	params := []map[string]interface{}{}

	if sp.Skill != "" {
		params = append(params, map[string]interface{}{"skill/role": Skill{Text: sp.Skill, Experience: "1-year-plus"}})
	}

	if sp.Organization != "" {
		params = append(params, map[string]interface{}{"organization": Term{Term: sp.Organization}})
	}

	if sp.Name != "" {
		params = append(params, map[string]interface{}{"name": Term{Term: sp.Name}})
	}

	if len(params) == 0 {
		return nil, nil
	}

	if len(params) > 1 {
		m["and"] = params
		bt, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}
		return bytes.NewBuffer(bt), nil
	}

	bt, err := json.Marshal(params[0])
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(bt), nil
}
