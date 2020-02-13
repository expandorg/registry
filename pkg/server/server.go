package server

import (
	"net/http"

	"github.com/expandorg/registry/pkg/authentication"

	"github.com/jmoiron/sqlx"

	"github.com/expandorg/registry/pkg/api/healthchecker"
	"github.com/expandorg/registry/pkg/api/registrationfetcher"
	"github.com/expandorg/registry/pkg/service"
	"github.com/gorilla/mux"
)

func New(
	db *sqlx.DB,
	s service.RegistryService,
) http.Handler {
	r := mux.NewRouter()

	r.Handle("/_health", healthchecker.MakeHandler(s)).Methods("GET")
	r.Handle("/registrations/{job_id}", registrationfetcher.MakeHandler(s)).Methods("GET")
	r.Use(authentication.AuthMiddleware)
	return withHandlers(r)
}
