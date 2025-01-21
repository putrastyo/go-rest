package middleware

import (
	"net/http"
	"praktik-todo/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get authorization from header
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Parse token
		token, err := utils.VerifyJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Set claims to ctx
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("email", claims["email"])
		}

		c.Next()
	}
}
