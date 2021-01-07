package data

// Strength Struct
type Strength struct {
	Name string `json:"name"`
}

// Interest Struct
type Interest struct {
	Name string `json:"name"`
}

// Organization struct
type Organization struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

// Experience struct
type Experience struct {
	ID               string         `json:"id"`
	Category         string         `json:"category"`
	Name             string         `json:"name"`
	Organizations    []Organization `json:"organizations"`
	Responsibilities []string       `json:"responsibilities"`
	FromMonth        string         `json:"fromMonth"`
	FromYear         string         `json:"fromYear"`
	ToMonth          string         `json:"toMonth"`
	ToYear           string         `json:"toYear"`
}

// Location struct
type Location struct {
	Name           string `json:"name"`
	Country        string `json:"country"`
	Timezone       string `json:"timezone"`
	TimezoneOffset int64  `json:"timezoneOffset"`
}

// Person struct
type Person struct {
	Name                 string   `json:"name"`
	Location             Location `json:"location"`
	OpenTo               []string `json:"openTo"`
	Picture              string   `json:"picture"`
	ProfessionalHeadline string   `json:"professionalHeadline"`
	// SubjectID            int64    `json:"subjectId"`
	Username string  `json:"username"`
	Verified bool    `json:"verified"`
	Weight   float64 `json:"weight"`
}

// BioData constains bio-data for a user
type BioData struct {
	Person       Person       `json:"person"`
	Strengths    []Strength   `json:"strengths"`
	Interests    []Interest   `json:"interests"`
	Experiences  []Experience `json:"experiences"`
	Awards       []Experience `json:"awards"`
	Jobs         []Experience `json:"jobs"`
	Projects     []Experience `json:"projects"`
	Publications []Experience `json:"publications"`
	Education    []Experience `json:"education"`
}
