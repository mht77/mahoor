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
	return sell, nil
}

func (s *sellRepository) DeleteSell(id uint) error {
	err := s.db.Delete(&models.Sell{}, id).Error
	return err
}
