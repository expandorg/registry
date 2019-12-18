package server

import (
	"net/http"

	"github.com/gemsorg/registry/pkg/authentication"

	"github.com/jmoiron/sqlx"

	"github.com/gemsorg/registry/pkg/api/healthchecker"
	"github.com/gemsorg/registry/pkg/service"
	"github.com/gorilla/mux"
)

func New(
	db *sqlx.DB,
	s service.RegistryService,
) http.Handler {
	r := mux.NewRouter()

	r.Handle("/_health", healthchecker.MakeHandler(s)).Methods("GET")
	r.Use(authentication.AuthMiddleware)
	return withHandlers(r)
}
