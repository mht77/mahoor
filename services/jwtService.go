package services

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/mht77/mahoor/models"
	"os"
	"time"
)

var JwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID     uint   `json:"user_id"`
	Username   string `json:"username"`
	IsAdmin    bool   `json:"is_admin"`
	IsApproved bool   `json:"is_approved"`
	jwt.RegisteredClaims
}

// GenerateToken creates a JWT for a user
func GenerateToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID:     user.Id,
		Username:   user.Username,
		IsAdmin:    user.IsAdmin,
		IsApproved: user.IsApproved,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}
