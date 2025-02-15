package contracts

type ProductCreationRequest struct {
	Name     string  `json:"name" binding:"required"`
	Quantity uint    `json:"quantity" binding:"required"`
	Price    float32 `json:"price" binding:"required"`
	TikkieId uint    `json:"tikkieId" binding:"required"`
}

type ProductUpdateRequest struct {
	Name     *string  `json:"name"`
	Quantity *uint    `json:"quantity"`
	Price    *float32 `json:"price"`
	TikkieId *uint    `json:"tikkieId"`
}
