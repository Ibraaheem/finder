package main

import (
	logging "finder/api/monitoring"
	"finder/api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	logger := logging.NewLogger()
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"} // You can restrict origins for security
	router.Use(cors.New(corsConfig))

	routes.SetupRoutes(router, logger)

	logger.Info("Server is running on :8080")
	router.Run(":8080")
}
