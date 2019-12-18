-include .env
LAST_MIGRATION = $(shell ls -tr migrations/sql/ | tail -n 1 | cut -d'_' -f1)

ifeq ($(LAST_MIGRATION),)
	LAST_MIGRATION := 0
endif


.PHONY: default

BIN_NAME=registry
VERSION ?= dev
GIT_COMMIT ?=$(shell git rev-parse HEAD)
SHORT_COMMIT ?=$(shell git rev-parse --short HEAD)
BUILD_DATE ?= $(shell date +%FT%T%z)
TIMESTAMP ?= $(shell date +%Y%m%d%H%M)

default: help

help:
	@echo 'Management commands for registry:'
	@echo
	@echo 'Usage:'
	@echo '    make build        			Builds the binary locally.'
	@echo '    make update-deps      	Runs dep ensure.'
	@echo '    make package       		Build final docker image with just the go binary inside.'
	@echo '    make add-migration  		Create a new migration file.'
	@echo '    make build-migration  	Create a new migration file.'
	@echo '    make run-migrations  	Run specific migration version and action.'
	@echo '    make migrate-latest  	Runs latest migration.'
	@echo '    make test          		Run tests on a compiled project.'
	@echo '    make run          			Build and run'
	@echo '    make up          			Start containers'
	@echo '    make down          		Stop and delete containers'
	@echo '    make deploy-dev    		Deploy tagged image to staging'
	@echo '    make deploy-prod   		Deploy tagged image to production'
	@echo '    make clean         		Clean the directory tree.'
	@echo


build: build-service

run: build 
	bin/registry

build-service:
	@echo "Building service"
	mkdir -p ./bin
	go build -ldflags "-w -X main.GitCommit=${GIT_COMMIT} -X main.Version=${VERSION} -X main.BuildDate=${BUILD_DATE}" -o ./bin/registry ./cmd/registry/

up:
	docker-compose up --build
	
update-deps:
	dep ensure -update

package: build-migrations
	@echo "Building image ${BIN_NAME} ${VERSION} $(GIT_COMMIT)"
	docker build --build-arg VERSION=${VERSION} --build-arg GIT_COMMIT=$(GIT_COMMIT) -t $(IMAGE_NAME):${VERSION} .

clean:
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}

deploy-dev: get-credentials-dev docker-build-dev push-dev
	kubectl set image deployment/registry registry=gcr.io/gems-org/registry-dev:$(VERSION)

deploy-prod: get-credentials-prod docker-build-prod push-prod
	kubectl set image deployment/registry registry=gcr.io/gems-org/registry:$(VERSION)

run-tests:
	go test ./... -v -count=1

down:
	docker-compose down

add-migration:
	touch migrations/sql/$(shell expr $(LAST_MIGRATION) + 1 )_$(name).up.sql
	touch migrations/sql/$(shell expr $(LAST_MIGRATION) + 1 )_$(name).down.sql
	echo $(shell expr $(LAST_MIGRATION) + 1 ) > migrations/version

build-migrations:
	docker build -t registry-migration migrations

run-migrations: build-migrations
	docker run --network host registry-migration \
	$(action) $(version) \
	"mysql://$(REGISTRY_DB_USER):$(REGISTRY_DB_PASSWORD)@tcp($(REGISTRY_DB_HOST):$(REGISTRY_DB_PORT))/$(REGISTRY_DB_NAME)"

migrate-latest: build-migrations
	docker run --network host registry-migration \
	goto $(LAST_MIGRATION) \
	"mysql://$(REGISTRY_DB_USER):$(REGISTRY_DB_PASSWORD)@tcp($(REGISTRY_DB_HOST):$(REGISTRY_DB_PORT))/$(REGISTRY_DB_NAME)"

db-seed:
	@echo "Seeding db"
	mkdir -p ./bin
	go build -o ./bin/dbseed ./pkg/database/dbseed
	./bin/dbseed

tag-staging:
	git tag -a staging-${SHORT_COMMIT} -m"staging release version ${SHORT_COMMIT}"