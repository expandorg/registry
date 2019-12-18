get-migrations-deps:
	go get -d -tags 'mysql file' github.com/golang-migrate/migrate/cmd/migrate

update-migrations-deps:
	go get -d -u -tags 'mysql file' github.com/golang-migrate/migrate/cmd/migrate

build-migrations:
	go build -o ./migrations/migrate -tags 'mysql file' github.com/golang-migrate/migrate/cmd/migrate