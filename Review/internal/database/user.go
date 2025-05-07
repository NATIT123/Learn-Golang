package database

import "context"

type User struct {
	ID   uint64
	Name string
}

//go:generate mockgen -source=$GOFILE -destination=$GOFILE.mock.go -package=$GOPACKAGE
type UserDataAccessor interface {
	GetUser(ctx context.Context, id uint64) (User, error)
	CreateUser(ctx context.Context, user User) error
	UpdateUser(ctx context.Context, user User) error
	DeleteUser(ctx context.Context, id uint64) error
}
