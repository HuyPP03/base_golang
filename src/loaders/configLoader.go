package loaders

import (
	"os"
	"strings"

	"github.com/HuyPP03/learn/src/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/unrolled/secure"
)

func ConfigLoader(router *gin.Engine) {
	// CORS Configuration
	corsOrigins := os.Getenv("CORS_ORIGINS")
	allowedOrigins := strings.Split(corsOrigins, ",")

	config := cors.Config{
		AllowOrigins:              allowedOrigins,
		AllowMethods:              []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:              []string{"Origin", "Content-Type", "X-Requested-With", "Accept", "Authorization", "access-token", "X-Access-Token"},
		AllowCredentials:          true,
		OptionsResponseStatusCode: 200,
	}
	router.Use(cors.New(config))

	secureLoader := secure.New(secure.Options{
		FrameDeny:          true,
		ContentTypeNosniff: true,
		ReferrerPolicy:     "no-referrer",
		SSLRedirect:        false,
		IsDevelopment:      gin.Mode() != gin.ReleaseMode,
	})
	router.Use(func(c *gin.Context) {
		err := secureLoader.Process(c.Writer, c.Request)
		if err != nil {
			return
		}
		c.Next()
	})

	router.Use(gin.LoggerWithWriter(logrus.StandardLogger().Out))
	router.Use(gin.Recovery())

	app := router.Group("/api")
	routes.APIRoutes(app)

	router.Static("/uploads", "./uploads")
}
