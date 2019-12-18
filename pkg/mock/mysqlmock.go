package mock

import (
	"database/sql"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func Mysql() (*sql.DB, *sqlx.DB, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return mockDB, sqlxDB, mock
}
