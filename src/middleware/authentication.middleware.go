package middleware

import (
	"net/http"

	"github.com/HuyPP03/learn/src/utils"
	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, utils.NewErrorResponse("Unauthorized", http.StatusUnauthorized))
			return
		}

		token := authHeader[7:]

		claims, err := utils.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.NewErrorResponse("Invalid token!", http.StatusUnauthorized))
			c.Abort()
			return
		}
		c.Set("user", claims)
		c.Next()
	}
}
