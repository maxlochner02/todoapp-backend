package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Kein Token vorhanden"})
			return
		}
		c.Set("user_id", 1) // Vereinfachung: Im echten Projekt Token parsen oder beim User-Service validieren
		c.Next()
	}
}
