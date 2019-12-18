package main

import (
	"fmt"
	"log"

	"github.com/gemsorg/registry/pkg/database"
	"github.com/jmoiron/sqlx"
	env "github.com/joho/godotenv"
)

func main() {
	err := env.Load()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal("mysql connection error", err)
	}
	defer db.Close()

	err = seed(db)
	if err != nil {
		log.Fatalln(err)
	}
}

func seed(db *sqlx.DB) error {
	fmt.Println("seeding db")
	return nil
}
