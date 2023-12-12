package models

import "gorm.io/gorm"

type Floors struct {
	gorm.Model
	ClassRoomName string
}

func (table *Floors) TableName() string {
	return "floors"
}
