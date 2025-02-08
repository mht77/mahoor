package models

import "time"

type User struct {
	Id         uint      `json:"id" gorm:"primary_key"`
	Username   string    `json:"username" gorm:"unique;not null"`
	Password   string    `json:"-" gorm:"not null"`
	IsAdmin    bool      `json:"isAdmin" gorm:"default:false"`
	IsApproved bool      `json:"isApproved" gorm:"default:false"`
	JoinedAt   time.Time `json:"joinedAt" gorm:"autoCreateTime"`
}
