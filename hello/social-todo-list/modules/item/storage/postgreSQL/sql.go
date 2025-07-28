package storage

import (
	"context"
	"log"
	model "main/modules/userlikeitem/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type sqlStore struct {
	db *gorm.DB
}

// Delete implements biz.UserUnlikeItemStore.
func (sql *sqlStore) Delete(ctx context.Context, userId int, itemId int) error {
	panic("unimplemented")
}

// Find implements biz.UserUnlikeItemStore.
func (sql *sqlStore) Find(ctx context.Context, userId int, itemId int) (*model.Like, error) {
	panic("unimplemented")
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
