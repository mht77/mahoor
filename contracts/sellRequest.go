package contracts

import (
	"time"
)

type SellCreationRequest struct {
	ProductId      uint    `form:"productId" binding:"required"`
	Quantity       *uint   `form:"quantity"`
	Name           *string `form:"name"`
	CollectionMode *string `form:"collectionMode"`
}

type SellResponse struct {
	Id uint `json:"id"`
	//Product   product.Product `json:"product"`
	Quantity       uint      `json:"quantity"`
	CreatedAt      time.Time `json:"createdAt"`
	Name           *string   `json:"name"`
	CollectionMode string    `json:"collectionMode"`
}
