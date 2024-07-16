package routes

import (
	"github.com/HuyPP03/learn/src/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {

	router := r.Group("/auth")

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.POST("/upload", controllers.Uploads)
}
