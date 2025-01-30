package models

import (
	"gorm.io/gorm"
	"time"
)

type Sell struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	ProductId uint           `json:"-"`
	Product   Product        `json:"product"`
	Quantity  uint           `json:"quantity"`
	CreateAt  time.Time      `json:"createAt" gorm:"autoCreateTime"`
	DeleteAt  gorm.DeletedAt `json:"-"`
}
