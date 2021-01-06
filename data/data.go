package data

// Data interface
type Data interface {
	GetBioDataByUsername(username string) (*BioData, error)
	GetJobByID(jobID string) (*Job, error)
}
