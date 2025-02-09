package repositories

import (
	"github.com/mht77/mahoor/models"
	"gorm.io/gorm"
)

type TikkieRepository interface {
	GetTikkies() (*[]models.Tikkie, error)
	CreateTikkie(*models.Tikkie) error
}

type tikkieRepo struct {
	db *gorm.DB
}

func NewTikkieRepository(db *gorm.DB) TikkieRepository {
	return &tikkieRepo{db: db}
}

func (t *tikkieRepo) GetTikkies() (*[]models.Tikkie, error) {
	var tikkies []models.Tikkie
	err := t.db.Find(&tikkies).Error
	if err != nil {
		return nil, err
	}
	return &tikkies, nil
}

func (t *tikkieRepo) CreateTikkie(tikkie *models.Tikkie) error {
	return t.db.Create(tikkie).Error
}
