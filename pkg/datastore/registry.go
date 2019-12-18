package datastore

import (
	"github.com/gemsorg/registry/pkg/registration"
	"github.com/jmoiron/sqlx"
)

type Storage interface {
	GetJobRegistrations(jobID uint64) (registration.Registrations, error)
}

type RegistryStore struct {
	DB *sqlx.DB
}

func NewRegistryStore(db *sqlx.DB) *RegistryStore {
	return &RegistryStore{
		DB: db,
	}
}

func (rs *RegistryStore) GetJobRegistrations(jobID uint64) (registration.Registrations, error) {
	reg := registration.Registrations{}

	err := rs.DB.Select(&reg, "SELECT * FROM registrations WHERE job_id=?", jobID)

	if err != nil {
		return reg, err
	}
	return reg, nil
}
