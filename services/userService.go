package services

import (
	"errors"
	"github.com/mht77/mahoor/contracts"
	"github.com/mht77/mahoor/models"
	"github.com/mht77/mahoor/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(userRequest *contracts.UserRequest) (*models.User, error)
	GetUserById(userId uint) (*models.User, error)
	GetAllUsers() (*[]models.User, error)
	GetJwt(userRequest *contracts.UserRequest) (string, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (u *userService) CreateUser(userRequest *contracts.UserRequest) (*models.User, error) {
	hashedPassword, err := HashPassword(userRequest.Password)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		Username: userRequest.Username,
		Password: hashedPassword,
	}
	return u.userRepository.CreateUser(user)
}

func (u *userService) GetUserById(userId uint) (*models.User, error) {
	return u.userRepository.GetUserById(userId)
}

func (u *userService) GetAllUsers() (*[]models.User, error) {
	return u.userRepository.GetAllUsers()
}

func (u *userService) GetJwt(userRequest *contracts.UserRequest) (string, error) {
	user, err := u.userRepository.GetUserByUsername(userRequest.Username)
	if err != nil {
		return "", errors.New("user not found")
	}
	if !CheckPassword(user.Password, userRequest.Password) {
		return "", errors.New("wrong password")
	}
	if !user.IsApproved {
		return "", errors.New("user is not approved")
	}
	jwt, err := GenerateToken(*user)
	if err != nil {
		return "", err
	}
	return jwt, nil
}

// HashPassword hashes a password using bcrypt.
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

// CheckPassword compares a hashed password with a plain one.
func CheckPassword(hashed, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}
