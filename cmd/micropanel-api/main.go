package main

import (
	"log"

	db "github.com/Turtel216/micro-panel/data"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	defer db.Close()
	log.Println("Successfully connected to database")

	//st := storer.NewMySQLStorer(db.GetDB())
	//str := server.NewServer(st)
}
