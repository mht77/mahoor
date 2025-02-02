package models

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID        uint           `json:"id" gorm:"primaryKey" gorm:"autoIncrement"`
	Name      string         `json:"name"`
	Quantity  uint           `json:"quantity"`
	Price     float32        `json:"price"`
	Available uint           `json:"available"`
	CreateAt  time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdateAt  time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Sells     []Sell         `json:"sells"`
}
