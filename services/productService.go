package services

import (
	"errors"
	"github.com/mht77/mahoor/contracts"
	"github.com/mht77/mahoor/models"
	"github.com/mht77/mahoor/repositories"
)

type ProductService interface {
	CreateProduct(product *contracts.ProductCreationRequest) (*models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	GetAllProducts() (*[]models.Product, error)
	UpdateProduct(id uint, product *contracts.ProductUpdateRequest) (*models.Product, error)
	DeleteProduct(id uint) error
}

type productService struct {
	productRepo repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (p *productService) CreateProduct(product *contracts.ProductCreationRequest) (*models.Product, error) {
	if product.Price < 0 {
		return nil, errors.New("price cannot be negative")
	}
	if product.Quantity < 0 {
		return nil, errors.New("quantity cannot be negative")
	}
	if len(product.Name) == 0 {
		return nil, errors.New("name cannot be empty")
	}
	return p.productRepo.CreateProduct(&models.Product{
		Name:      product.Name,
		Quantity:  product.Quantity,
		Price:     product.Price,
		Available: product.Quantity,
		TikkieId:  product.TikkieId,
	})
}

func (p *productService) GetProductByID(id uint) (*models.Product, error) {
	return p.productRepo.GetProductByID(id)
}

func (p *productService) GetAllProducts() (*[]models.Product, error) {
	return p.productRepo.GetAllProducts()
}

func (p *productService) UpdateProduct(id uint, product *contracts.ProductUpdateRequest) (*models.Product, error) {
	if product.Price != nil && *product.Price < 0 {
		return nil, errors.New("price cannot be negative")
	}
	if product.Quantity != nil && *product.Quantity < 0 {
		return nil, errors.New("quantity cannot be negative")
	}
	if product.Name != nil && len(*product.Name) == 0 {
		return nil, errors.New("name cannot be empty")
	}
	oldProduct, err := p.productRepo.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	sellsCount := oldProduct.Quantity - oldProduct.Available
	if product.Name == nil {
		product.Name = &oldProduct.Name
	}
	if product.Quantity == nil {
		product.Quantity = &oldProduct.Quantity
	}
	if product.Price == nil {
		product.Price = &oldProduct.Price
	}
	return p.productRepo.UpdateProduct(&models.Product{
		Id:        id,
		Name:      *product.Name,
		Quantity:  *product.Quantity,
		Price:     *product.Price,
		Sells:     oldProduct.Sells,
		Available: oldProduct.Available + sellsCount,
		TikkieId:  oldProduct.TikkieId,
	})
}

func (p *productService) DeleteProduct(id uint) error {
	return p.productRepo.DeleteProduct(id)
}
