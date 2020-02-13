package test

import (
	"database/sql"
	"os"
	"testing"

	"github.com/expandorg/registry/pkg/mock"
	"github.com/expandorg/registry/pkg/service"
	"github.com/golang/mock/gomock"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

var jwtToken, _ = mock.GenerateJWT(8)

var bearer = "Bearer " + jwtToken

func Setup(t *testing.T) (*sql.DB, *sqlx.DB, sqlmock.Sqlmock, *service.MockRegistryService) {
	os.Setenv("JWT_SECRET", mock.JWT_SECRET)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	s := service.NewMockRegistryService(ctrl)
	db, dbx, m := mock.Mysql()
	return db, dbx, m, s
}
