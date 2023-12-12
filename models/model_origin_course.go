package models

import (
	"gorm.io/gorm"
)

type OriginClass struct {
	gorm.Model
	TeacherName string
	TeacherRoom string
	CourseName  string
	ClassName   string
	Population  string
	Software    string
	Cycle       string
}

func (table *OriginClass) TableName() string {
	return "origin_class"
}
