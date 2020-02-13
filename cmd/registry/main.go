package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/expandorg/registry/pkg/authorization"
	"github.com/expandorg/registry/pkg/database"
	"github.com/expandorg/registry/pkg/datastore"
	"github.com/expandorg/registry/pkg/service"
	"github.com/joho/godotenv"

	"github.com/expandorg/registry/pkg/server"
)

func main() {
	environment := flag.String("env", "local", "use compose in compose-dev")
	flag.Parse()

	if *environment == "local" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// Connect to db
	db, err := database.Connect()
	if err != nil {
		log.Fatal("mysql connection error", err)
	}
	defer db.Close()
	ds := datastore.NewRegistryStore(db)
	authorizer := authorization.NewAuthorizer()
	svc := service.New(ds, authorizer)
	s := server.New(db, svc)
	log.Println("info", fmt.Sprintf("Starting service on port 8184"))
	http.Handle("/", s)
	http.ListenAndServe(":8184", nil)
}
