package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mht77/mahoor/contracts"
	"github.com/mht77/mahoor/services"
)

type TikkieController struct {
	tikkieService services.TikkieService
}

func NewTikkieController(s services.TikkieService) *TikkieController {
	return &TikkieController{
		tikkieService: s,
	}
}

// GetTikkies godoc
// @Summary Get tikkies
// @Description Get all tikkies
// @Tags tikkies
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Tikkie
// @Failure 400 {object} string
// @Failure 401 {object} string
// @Router /tikkies [get]
func (tc *TikkieController) GetTikkies(c *gin.Context) {
	tikkies, err := tc.tikkieService.GetTikkies()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, tikkies)
}

// CreateTikkie godoc
// @Summary Create a tikkie
// @Description Create a tikkie
// @Tags tikkies
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param tikkieRequest body contracts.TikkieRequest true "Tikkie creation request"
// @Success 204
// @Failure 400 {object} string
// @Failure 401 {object} string
// @Router /tikkies [post]
func (tc *TikkieController) CreateTikkie(c *gin.Context) {
	var tikkieRequest contracts.TikkieRequest
	err := c.BindJSON(&tikkieRequest)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = tc.tikkieService.CreateTikkie(&tikkieRequest)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, nil)
}
