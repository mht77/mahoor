package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mht77/mahoor/contracts"
	_ "github.com/mht77/mahoor/models"
	"github.com/mht77/mahoor/services"
	"strconv"
)

type SellController struct {
	sellService services.SellService
}

func NewSellController(sellService services.SellService) *SellController {
	return &SellController{
		sellService: sellService,
	}
}

// CreateSell godoc
// @Summary Create a sell
// @Description Create a sell
// @Tags sells
// @Accept json
// @Produce json
// @Param productId query int true "Product ID"
// @Param quantity query int false "Quantity"
// @Success 201 {object} models.Sell
// @Failure 400 {object} string
// @Router /sells [get]
func (s *SellController) CreateSell(c *gin.Context) {
	var sellCreationRequest contracts.SellCreationRequest
	err := c.ShouldBindQuery(&sellCreationRequest)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	sell, err := s.sellService.CreateSell(&sellCreationRequest)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, sell)
}

// GetSellsByProductID godoc
// @Summary Get sells by product ID
// @Description Get sells by product ID
// @Tags sells
// @Accept json
// @Produce json
// @Param productId path int true "Product ID"
// @Success 200 {array} models.Sell
// @Failure 400 {object} string
// @Router /sells/{productId} [get]
func (s *SellController) GetSellsByProductID(c *gin.Context) {
	productId, err := strconv.ParseUint(c.Param("productId"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid product id"})
		return
	}
	sells, err := s.sellService.GetSellsByProductID(uint(productId))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, sells)
}
