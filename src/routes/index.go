package routes

import "github.com/gin-gonic/gin"

func APIRoutes(r *gin.RouterGroup) {
	AuthRoutes(r)
	UserRoutes(r)
}
