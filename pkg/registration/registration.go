package registration

import "encoding/json"

type Registration struct {
	ID          uint64          `json:"id" db:"id"`
	JobID       uint64          `json:"job_id" db:"job_id"`
	APIKeyID    string          `json:"api_key_id" db:"api_key_id"`
	Services    json.RawMessage `json:"services" db:"services"`
	RequesterID uint64          `json:"requester_id" db:"requester_id"`
}
