package main

import (
	"log"

	db "github.com/Turtel216/micro-panel/data"
	"github.com/Turtel216/micro-panel/micropanel-api/handler"
	"github.com/Turtel216/micro-panel/micropanel-api/server"
	"github.com/Turtel216/micro-panel/micropanel-api/storer"
	"github.com/ianschenck/envflag"
)

const minSecretKeySize = 32

func main() {
	var secretKey = envflag.String("SECRET_KEY", "01234567890123456789012345678901", "secret key for JWT signing")
	if len(*secretKey) < minSecretKeySize {
		log.Fatalf("SECRET_KEY must be at least %d characters long", minSecretKeySize)
	}

	db, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	defer db.Close()
	log.Println("Successfully connected to database")

	st := storer.NewMySQLStorer(db.GetDB())
	srv := server.NewServer(st)
	hdl := handler.NewHandler(srv, *secretKey)
	handler.RegisterRoutes(hdl)
	handler.Start(":8080")
}
