package common

import "main/modules/user/models/enum"

type SimpleUser struct {
	SQLModel
	LastName  string           `json:"last_name" gorm:"column:last_name;"`
	FirstName string           `json:"first_name" gorm:"column:first_name;"`
	Status    *enum.UserStatus `json:"status" gorm:"column:status;"`
	Role      *enum.UserRole   `json:"role" gorm:"column:role;"`
}

func (SimpleUser) TableName() string {
	return "users"
}

func (u *SimpleUser) Mask() {
	u.SQLModel.Mask(DbTypeUser)
}
