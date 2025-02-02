package repositories

import (
	"github.com/mht77/mahoor/models"
	"gorm.io/gorm"
)

type SellRepository interface {
	CreateSell(sell *models.Sell) (*models.Sell, error)
	DeleteSell(id uint) error
	GetSellsByProductID(productId uint) ([]models.Sell, error)
	GetAllSells() ([]models.Sell, error)
}

type sellRepository struct {
	db *gorm.DB
}

func NewSellRepository(db *gorm.DB) SellRepository {
	return &sellRepository{
		db: db,
	}
}

func (s *sellRepository) GetSellsByProductID(productId uint) ([]models.Sell, error) {
	var sells []models.Sell
	err := s.db.Preload("Product").Where("product_id = ?", productId).Find(&sells).Error
	if err != nil {
		return nil, err
	}
	return sells, nil
}

func (s *sellRepository) GetAllSells() ([]models.Sell, error) {
	var sells []models.Sell
	err := s.db.Preload("Product").Find(&sells).Error
	if err != nil {
		return nil, err
	}
	return sells, nil
}

func (s *sellRepository) CreateSell(sell *models.Sell) (*models.Sell, error) {
	err := s.db.Create(&sell).Error
	if err != nil {
		return nil, err
	}
	var product models.Product
	s.db.Model(&models.Product{}).Where("id = ?", sell.ProductId).Find(&product)
	product.Available -= sell.Quantity
	s.db.Save(&product)
	sell.Product = product
	return sell, nil
}

func (s *sellRepository) DeleteSell(id uint) error {
	var sell models.Sell
	err := s.db.First(&sell, id).Error
	if err != nil {
		return err
	}
	var product models.Product
	err = s.db.Model(&models.Product{}).Where("id = ?", sell.ProductId).Find(&product).Error
	if err != nil {
		return err
	}
	product.Available += sell.Quantity
	s.db.Save(&product)
	s.db.Delete(&models.Sell{}, id)
	return err
}
