package models

import "gorm.io/gorm"

type Software struct {
	gorm.Model
	Software      string
	ClassRoomName string
}

func (table *Software) TableName() string {
	return "software"
}
