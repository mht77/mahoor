package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mht77/mahoor/contracts"
	_ "github.com/mht77/mahoor/models"
	"github.com/mht77/mahoor/services"
	"strconv"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

// CreateProduct godoc
// @Summary Create a product
// @Description Create a product
// @Tags products
// @Accept json
// @Produce json
// @Param productCreationRequest body contracts.ProductCreationRequest true "Product Creation Request"
// @Success 201 {object} models.Product
// @Failure 400 {object} string
// @Router /products [post]
func (p *ProductController) CreateProduct(c *gin.Context) {
	var productCreationRequest contracts.ProductCreationRequest
	if err := c.ShouldBindJSON(&productCreationRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	product, err := p.productService.CreateProduct(&productCreationRequest)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, product)
}

// GetProductByID godoc
// @Summary Get a product by Id
// @Description Get a product by Id
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product Id"
// @Success 200 {object} models.Product
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /products/{id} [get]
func (p *ProductController) GetProductByID(c *gin.Context) {
	// parse the id from the request to uint
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	product, err := p.productService.GetProductByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, product)
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Get all products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Failure 500 {object} string
// @Router /products [get]
func (p *ProductController) GetAllProducts(c *gin.Context) {
	products, err := p.productService.GetAllProducts()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, products)

}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product Id"
// @Param productUpdateRequest body contracts.ProductUpdateRequest true "Product Update Request"
// @Success 200 {object} models.Product
// @Failure 400 {object} string
// @Router /products/{id} [put]
func (p *ProductController) UpdateProduct(c *gin.Context) {
	var productUpdateRequest contracts.ProductUpdateRequest
	if err := c.ShouldBindJSON(&productUpdateRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	product, err := p.productService.UpdateProduct(uint(id), &productUpdateRequest)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, product)
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product Id"
// @Success 204
// @Failure 400 {object} string
// @Router /products/{id} [delete]
func (p *ProductController) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	err = p.productService.DeleteProduct(uint(id))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, nil)
}
