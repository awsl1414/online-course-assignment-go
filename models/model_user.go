package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName       string
	JobNumber      string
	HashedPassword string
	Permission     string
}

func (table *User) TableName() string {
	return "user"
}
