package storage

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type sqlStore struct {
	db *gorm.DB
}

func CreateSQL(DB_CONN_STR string) *gorm.DB {
	// https://github.com/jackc/pgx
	///Connect PostgreSQL
	db, err := gorm.Open(postgres.Open(DB_CONN_STR), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
