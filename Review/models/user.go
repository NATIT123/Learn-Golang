package models

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	ID    uint    `gorm:"primaryKey"`
	Name  string  `json:"name"`
	Email *string `json:"email" gorm:"unique"`
}

type UserGorm struct {
	gorm.Model
	Name string
}

type UserStorage interface {
	GetAllUsers(ctx context.Context) ([]User, error)
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, id uint) (User, error)
	UpdateUser(ctx context.Context, user User) error
	DeleteUser(ctx context.Context, id uint) error
}

type gormUserStorage struct {
	DB *gorm.DB
}

func NewUserStorage(db *gorm.DB) UserStorage {
	return &gormUserStorage{DB: db}
}

func (s *gormUserStorage) GetAllUsers(ctx context.Context) ([]User, error) {
	var users []UserGorm
	if err := s.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	var result []User
	for _, u := range users {
		result = append(result, User{ID: u.ID, Name: u.Name})
	}
	return result, nil
}

func (s *gormUserStorage) CreateUser(ctx context.Context, user User) error {
	if err := s.DB.Create(&UserGorm{
		Name: user.Name,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (s *gormUserStorage) GetUser(ctx context.Context, id uint) (User, error) {
	var user UserGorm
	if err := s.DB.First(&user, id).Error; err != nil {
		return User{}, err
	}
	return User{ID: user.ID, Name: user.Name}, nil
}

func (s *gormUserStorage) UpdateUser(ctx context.Context, user User) error {
	if err := s.DB.Model(&UserGorm{}).Where("id = ?", user.ID).Updates(UserGorm{Name: user.Name}).Error; err != nil {
		return err
	}
	return nil
}

func (s *gormUserStorage) DeleteUser(ctx context.Context, id uint) error {
	if err := s.DB.Delete(&UserGorm{}, id).Error; err != nil {
		return err
	}
	return nil
}
