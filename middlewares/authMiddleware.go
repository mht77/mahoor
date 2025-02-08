package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mht77/mahoor/services"
	"strings"
)

// AuthMiddleware checks if a request has a valid JWT token.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// Ensure token is provided in the correct format
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(401, gin.H{"error": "Unauthorized: Missing token"})
			c.Abort()
			return
		}

		// Extract token value
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims := &services.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return services.JwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(403, gin.H{"error": "Unauthorized: Invalid token"})
			c.Abort()
			return
		}

		// Store user info in context
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("is_admin", claims.IsAdmin)
		c.Set("is_approved", claims.IsApproved)

		c.Next()
	}
}
