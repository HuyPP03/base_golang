package routes

import (
	"github.com/HuyPP03/learn/src/controllers"
	"github.com/HuyPP03/learn/src/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	router := r.Group("/user")
	router.Use(middleware.AuthenticationMiddleware())
	router.Use(middleware.AuthorizationMiddleware())
	router.POST("/profile", controllers.GetProfile)
}
