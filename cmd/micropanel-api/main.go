package main

import (
	"log"

	db "github.com/Turtel216/micro-panel/data"
	"github.com/Turtel216/micro-panel/micropanel-api/handler"
	"github.com/Turtel216/micro-panel/micropanel-api/server"
	"github.com/Turtel216/micro-panel/micropanel-api/storer"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	defer db.Close()
	log.Println("Successfully connected to database")

	st := storer.NewMySQLStorer(db.GetDB())
	srv := server.NewServer(st)
	hdl := handler.NewHandler(srv)
	handler.RegisterRoutes(hdl)
	handler.Start(":8080")
}
