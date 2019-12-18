package registration

type Registration struct {
	ID          uint64 `json:"id" db:"id"`
	URL         string `json:"url" db:"url"`
	Service     string `json:"service" db:"service"`
	JobID       uint64 `json:"job_id" db:"job_id"`
	RequesterID uint64 `json:"requester_id" db:"requester_id"`
	Active      bool   `json:"active" db:"active"`
}

type Registrations []Registration
