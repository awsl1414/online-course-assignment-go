package models

import "gorm.io/gorm"

type DetailClass struct {
	gorm.Model
	TeacherName      string
	TeacherRoom      string
	CourseName       string
	ClassName        string
	Population       string
	Software         string
	ComputerRoomName string
	Day              string
	Lesson           string
	Cycle            string
}

func (table *DetailClass) TableName() string {
	return "detail_class"
}
