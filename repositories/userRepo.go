package repositories

import (
	"github.com/mht77/mahoor/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserById(id uint) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetAllUsers() (*[]models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) CreateUser(user *models.User) (*models.User, error) {
	err := u.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) GetUserById(id uint) (*models.User, error) {
	var user models.User
	err := u.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := u.db.First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) GetAllUsers() (*[]models.User, error) {
	var users []models.User
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}
