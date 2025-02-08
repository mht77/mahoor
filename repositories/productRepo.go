package repositories

import (
	"github.com/mht77/mahoor/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) (*models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	GetAllProducts() (*[]models.Product, error)
	UpdateProduct(product *models.Product) (*models.Product, error)
	DeleteProduct(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (p *productRepository) CreateProduct(product *models.Product) (*models.Product, error) {
	err := p.db.Create(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *productRepository) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	err := p.db.Preload("Sells").First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *productRepository) GetAllProducts() (*[]models.Product, error) {
	var products []models.Product
	err := p.db.Preload("Sells").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return &products, nil
}

func (p *productRepository) UpdateProduct(product *models.Product) (*models.Product, error) {
	err := p.db.Updates(product).Error
	if err != nil {
		return nil, err
	}
	var updatedProduct models.Product
	err = p.db.First(&updatedProduct, product.Id).Error
	if err != nil {
		return nil, err
	}
	return &updatedProduct, nil
}

func (p *productRepository) DeleteProduct(id uint) error {
	p.db.Delete(&models.Product{}, id)
	p.db.Model(&models.Sell{}).Where("product_id = ?", id).Delete(&models.Sell{})
	return nil
}
