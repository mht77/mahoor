package services

import (
	"github.com/mht77/mahoor/contracts"
	"github.com/mht77/mahoor/models"
	"github.com/mht77/mahoor/repositories"
)

type SellService interface {
	CreateSell(sellRequest *contracts.SellCreationRequest) (*models.Sell, error)
	DeleteSell(id uint) error
	GetSellsByProductID(productId uint) (*[]models.Sell, error)
	GetAllSells() (*[]models.Sell, error)
}

type sellService struct {
	sellRepo repositories.SellRepository
}

func NewSellService(sellRepo repositories.SellRepository) SellService {
	return &sellService{
		sellRepo: sellRepo,
	}
}

func (s *sellService) CreateSell(sellRequest *contracts.SellCreationRequest) (*models.Sell, error) {
	sell := &models.Sell{
		ProductId:      sellRequest.ProductId,
		Quantity:       1,
		Name:           sellRequest.Name,
		CollectionMode: models.EatIn,
	}
	if sellRequest.Quantity != nil && *sellRequest.Quantity > 0 {
		sell.Quantity = *sellRequest.Quantity
	}
	if sellRequest.CollectionMode != nil && *sellRequest.CollectionMode != "" {
		mode := models.CollectionMode(*sellRequest.CollectionMode)
		if mode == models.Takeaway {
			sell.CollectionMode = mode
		}
	}
	return s.sellRepo.CreateSell(sell)
}

func (s *sellService) DeleteSell(id uint) error {
	return s.sellRepo.DeleteSell(id)
}

func (s *sellService) GetSellsByProductID(productId uint) (*[]models.Sell, error) {
	return s.sellRepo.GetSellsByProductID(productId)
}

func (s *sellService) GetAllSells() (*[]models.Sell, error) {
	return s.sellRepo.GetAllSells()
}
