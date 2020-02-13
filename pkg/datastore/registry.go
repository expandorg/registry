package datastore

import (
	"github.com/expandorg/registry/pkg/registration"
	"github.com/jmoiron/sqlx"
)

type Storage interface {
	GetJobRegistration(jobID uint64) (registration.Registration, error)
}

type RegistryStore struct {
	DB *sqlx.DB
}

func NewRegistryStore(db *sqlx.DB) *RegistryStore {
	return &RegistryStore{
		DB: db,
	}
}

func (rs *RegistryStore) GetJobRegistration(jobID uint64) (registration.Registration, error) {
	reg := []registration.Registration{}

	err := rs.DB.Select(&reg, "SELECT * FROM registrations WHERE job_id=?", jobID)

	if err != nil {
		return registration.Registration{}, err
	}

	if len(reg) == 0 {
		return registration.Registration{}, nil
	}

	return reg[0], nil
}
