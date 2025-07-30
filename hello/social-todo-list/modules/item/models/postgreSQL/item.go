package models

import (
	"errors"
	"main/common"
)

const (
	EntityName = "Item"
)

var (
	ErrTitleIsBlank = errors.New("title can not be blank")
	ErrItemDeleted  = errors.New("item is deleted")
)

type TodoItem struct {
	common.SQLModel
	UserId      int                `json:"-" gorm:"column:user_id"`
	Title       string             `json:"title" gorm:"column:title;"`
	Description string             `json:"description" gorm:"column:description;"`
	Status      *ItemStatus        `json:"status" gorm:"column:status;"`
	Image       *common.Image      `json:"image" gorm:"column:image;"`
	Cover       *common.Images     `json:"cover" gorm:"column:cover;"`
	LikedCount  int                `json:"liked_count" gorm:"-"`
	Owner       *common.SimpleUser `json:"owner" gorm:"foreignKey:UserId;"`
}

func (i *TodoItem) Mask() {
	i.SQLModel.Mask(common.DbTypeItem)

	if v := i.Owner; v != nil {
		v.Mask()
	}
}

func (TodoItem) TableName() string { return "todo_items" }

type TodoItemCreation struct {
	Id          int            `json:"id" gorm:"column:id;"`
	UserId      int            `json:"-" gorm:"column:user_id;"`
	Title       string         `json:"title" gorm:"column:title;"`
	Description string         `json:"description" gorm:"column:description"`
	Status      *ItemStatus    `json:"status" gorm:"column:status"`
	Cover       *common.Images `json:"cover" gorm:"column:cover;"`
	Image       *common.Image  `json:"image" gorm:"column:image;"`
}

type TodoItemUpdate struct {
	Title       *string        `json:"title" gorm:"column:title;"`
	Description string         `json:"description" gorm:"column:description"`
	Image       *common.Image  `json:"image" gorm:"column:image;"`
	Cover       *common.Images `json:"cover" gorm:"column:cover;"`
	Status      *ItemStatus    `json:"status" gorm:"column:status;"`
}

func (TodoItemUpdate) TableName() string { return TodoItem{}.TableName() }

func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }
