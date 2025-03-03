package models

import "gorm.io/gorm"

type Attendance struct {
	Id        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string         `json:"name"`
	Number    uint           `json:"number"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
