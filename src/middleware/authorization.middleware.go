package middleware

import (
	"net/http"

	"github.com/HuyPP03/learn/src/utils"
	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, utils.NewErrorResponse("Unauthorized", http.StatusUnauthorized))
			c.Abort()
			return
		}
		claims := user.(*utils.Claims)
		if claims.Role != "admin" {
			c.JSON(http.StatusUnauthorized, utils.NewErrorResponse("Unauthorized", http.StatusUnauthorized))
			c.Abort()
			return
		}

		c.Next()
	}
}
