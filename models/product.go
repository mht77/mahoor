package models

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	Id                uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name              string         `json:"name"`
	Quantity          uint           `json:"quantity"`
	Price             float32        `json:"price"`
	Available         uint           `json:"available"`
	ExcludeInPreorder bool           `json:"excludeInPreorder" default:"false"`
	CreateAt          time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdateAt          time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt         gorm.DeletedAt `json:"-"`
	Sells             *[]Sell        `json:"sells"`
	Tikkie            Tikkie         `json:"tikkie"`
	TikkieId          uint           `json:"-"`
}
