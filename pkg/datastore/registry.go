package datastore

import (
	"github.com/jmoiron/sqlx"
)

type Storage interface {
}

type RegistryStore struct {
	DB *sqlx.DB
}

func NewRegistryStore(db *sqlx.DB) *RegistryStore {
	return &RegistryStore{
		DB: db,
	}
}
