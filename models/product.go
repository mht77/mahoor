package models

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Quantity  uint           `json:"quantity"`
	Price     float32        `json:"price"`
	CreateAt  time.Time      `json:"createAt" gorm:"autoCreateTime"`
	UpdateAt  time.Time      `json:"updateAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Sells     []Sell         `json:"sells"`
}
