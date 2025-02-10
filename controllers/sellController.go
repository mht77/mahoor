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
// @Summary Create Sell
// @Description Create Sell by providing product id and optional quantity
// @Tags sells
// @Accept json
// @Produce json
// @Param sellCreationRequest body contracts.SellCreationRequest true "Sell Creation Request"
// @Success 201 {object} models.Sell
// @Failure 400 {object} string
// @Router /sells/ [post]
func (s *SellController) CreateSell(c *gin.Context) {
	var sellCreationRequest contracts.SellCreationRequest
	err := c.ShouldBindJSON(&sellCreationRequest)
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
// @Summary Get sells by product Id
// @Description Get sells by product Id
// @Tags sells
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param productId path int true "Product Id"
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

// GetAllSells godoc
// @Summary Get all sells
// @Description Get all sells
// @Tags sells
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Sell
// @Failure 500 {object} string
// @Router /sells/ [get]
func (s *SellController) GetAllSells(c *gin.Context) {
	sells, err := s.sellService.GetAllSells()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, sells)
}

// DeleteSell godoc
// @Summary Delete a sell
// @Description Delete a sell by id
// @Tags sells
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 204
// @Failure 400 {object} string
// @Router /sells/{id} [delete]
func (s *SellController) DeleteSell(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	err = s.sellService.DeleteSell(uint(id))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, nil)
}
