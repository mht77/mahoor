package contracts

import (
	"time"
)

type SellCreationRequest struct {
	ProductId uint  `form:"productId" binding:"required"`
	Quantity  *uint `form:"quantity"`
}

type SellResponse struct {
	Id uint `json:"id"`
	//Product   product.Product `json:"product"`
	Quantity  uint      `json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
}
