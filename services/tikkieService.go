package services

import (
	"github.com/mht77/mahoor/contracts"
	"github.com/mht77/mahoor/models"
	"github.com/mht77/mahoor/repositories"
)

type TikkieService interface {
	GetTikkies() (*[]models.Tikkie, error)
	CreateTikkie(tikkieRequest *contracts.TikkieRequest) error
}

type tikkieService struct {
	tikkieRepository repositories.TikkieRepository
}

func NewTikkieService(tikkieRepository repositories.TikkieRepository) TikkieService {
	return &tikkieService{tikkieRepository: tikkieRepository}
}

func (t *tikkieService) GetTikkies() (*[]models.Tikkie, error) {
	return t.tikkieRepository.GetTikkies()
}

func (t *tikkieService) CreateTikkie(tikkieRequest *contracts.TikkieRequest) error {
	tikkie := models.Tikkie{
		Nickname: tikkieRequest.Nickname,
		Link:     tikkieRequest.Link,
	}
	return t.tikkieRepository.CreateTikkie(&tikkie)
}
