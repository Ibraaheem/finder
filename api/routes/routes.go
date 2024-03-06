package routes

import (
	"finder/api/handlers"
	logging "finder/api/monitoring"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, logger *logging.Logger) {
	router.GET("/health", func(c *gin.Context) {
		handlers.HealthCheckHandler(c)
	})

	router.GET("/explore", func(c *gin.Context) {
		handlers.ExploreHandler(c, logger)
	})
}
