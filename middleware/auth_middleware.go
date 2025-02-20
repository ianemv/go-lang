package middleware

import (
	"backend/utils"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Debug logging
		log.Printf("Starting auth middleware check...")

		authHeader := c.GetHeader("Authorization")
		log.Printf("Auth header received: %s", authHeader)

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		token := strings.Split(authHeader, "Bearer ")
		if len(token) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(token[1])
		if err != nil {
			log.Printf("Token validation error: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		log.Printf("Token validated successfully, UserID: %v", claims.UserID)

		// Set user_id in context
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
