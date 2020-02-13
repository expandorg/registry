package service

import (
	"github.com/expandorg/registry/pkg/authentication"
	"github.com/expandorg/registry/pkg/authorization"
	"github.com/expandorg/registry/pkg/datastore"
	"github.com/expandorg/registry/pkg/registration"
)

type RegistryService interface {
	Healthy() bool
	SetAuthData(data authentication.AuthData)
	GetJobRegistration(jobID uint64) (registration.Registration, error)
}

type service struct {
	store      datastore.Storage
	authorizor authorization.Authorizer
}

func New(s datastore.Storage, a authorization.Authorizer) *service {
	return &service{
		store:      s,
		authorizor: a,
	}
}

func (s *service) Healthy() bool {
	return true
}

func (s *service) SetAuthData(data authentication.AuthData) {
	s.authorizor.SetAuthData(data)
}

func (s *service) GetJobRegistration(jobID uint64) (registration.Registration, error) {
	return s.store.GetJobRegistration(jobID)
}
