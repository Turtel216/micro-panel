package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	db *sqlx.DB
}

func NewDatabase() (*DB, error) {
	db, err := sqlx.Open("mysql", "root:passowrd@tcp(localhost:3306)/micropanel?parseTime=true")
	if err != nil {
		return nil, fmt.Errorf("error opening dabase: %w", err)
	}

	return &DB{db: db}, err
}

func (d *DB) Close() error {
	return d.db.Close()
}

func (d *DB) GetDB() *sqlx.DB {
	return d.db
}
