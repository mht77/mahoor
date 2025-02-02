package models

import (
	"gorm.io/gorm"
	"time"
)

type Sell struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	Name      *string        `json:"name"`
	ProductId uint           `json:"-"`
	Product   Product        `json:"product"`
	Quantity  uint           `json:"quantity"`
	CreateAt  time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	DeleteAt  gorm.DeletedAt `json:"-"`
}
