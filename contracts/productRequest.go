package contracts

import "mime/multipart"

type ProductCreationRequest struct {
	Name              string  `form:"name" binding:"required"`
	Quantity          uint    `form:"quantity" binding:"required"`
	Price             float32 `form:"price" binding:"required"`
	TikkieId          uint    `form:"tikkieId" binding:"required"`
	ExcludeInPreorder bool    `form:"excludeInPreorder" default:"false"`
	StopPreorderAt    int     `form:"stopPreorderAt" default:"0"`
	Picture           *string
	PictureFile       *multipart.FileHeader `form:"pictureFile"`
}

type ProductUpdateRequest struct {
	Name              *string  `form:"name"`
	Quantity          *uint    `form:"quantity"`
	Price             *float32 `form:"price"`
	TikkieId          *uint    `form:"tikkieId"`
	ExcludeInPreorder *bool    `form:"excludeInPreorder"`
	StopPreorderAt    *int     `form:"stopPreorderAt"`
	Picture           *string
	PictureFile       *multipart.FileHeader `form:"pictureFile"`
}
