package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mht77/mahoor/contracts"
	"github.com/mht77/mahoor/services"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// CreateUser godoc
// @Summary sign up
// @Description sign up
// @Tags users
// @Accept json
// @Produce json
// @Param userRequest body contracts.UserRequest true "User signup request"
// @Success 201 {object} models.User
// @Failure 400 {object} string
// @Router /users [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	var userRequest contracts.UserRequest
	err := c.BindJSON(&userRequest)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err := uc.userService.CreateUser(&userRequest)
	if err != nil {
		c.JSON(400, gin.H{"error": "Username already exists"})
		return
	}
	c.JSON(201, user)
}

// GetAllUsers godoc
// @Summary Get users
// @Description Get All Users
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.User
// @Failure 400 {object} string
// @Failure 401 {object} string
// @Router /users [get]
func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}

// GetToken godoc
// @Summary Get token
// @Description Get token by providing credentials
// @Tags users
// @Accept json
// @Produce json
// @Param userRequest body contracts.UserRequest true "User signup request"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /users/token [post]
func (uc *UserController) GetToken(c *gin.Context) {
	var userRequest contracts.UserRequest
	err := c.BindJSON(&userRequest)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk, err := uc.userService.GetJwt(&userRequest)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": tk})
}
