package models

import (
	"errors"
	"main/common"
	"main/modules/user/models/enum"
)

const (
	EntityName = "User"
)

type User struct {
	common.SQLModel
	Email     string           `json:"email" gorm:"column:email;"`
	Password  string           `json:"-" gorm:"column:password;"`
	Salt      string           `json:"-" gorm:"column:salt;"`
	LastName  string           `json:"last_name" gorm:"column:last_name;"`
	FirstName string           `json:"first_name" gorm:"column:first_name;"`
	Phone     string           `json:"phone" gorm:"column:phone;"`
	Status    *enum.UserStatus `json:"status" gorm:"column:status;"`
	Role      *enum.UserRole   `json:"role" gorm:"column:role;"`
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role.String()
}

func (User) TableName() string {
	return "users"
}

type UserCreate struct {
	common.SQLModel `json:"-" gorm:"inline"`
	Email           string `json:"email" gorm:"column:email;"`
	Password        string `json:"password" gorm:"column:password;"`
	LastName        string `json:"last_name" gorm:"column:last_name;"`
	FirstName       string `json:"first_name" gorm:"column:first_name;"`
	Role            string `json:"role" gorm:"column:role;"`
	Phone           string `json:"phone" gorm:"column:phone;"`
	Salt            string `json:"-" gorm:"column:salt;"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"column:password;"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

var (
	ErrEmailOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)
	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)

	ErrUserDeleted = common.NewCustomError(
		errors.New("user has been deleted"),
		"user has been deleted",
		"ErrUserDeleted",
	)
)
