package handlers

import (
	"finder/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(c *gin.Context) {
	healthStatus := gin.H{"status": "healthy"}

	response := models.SuccessResponse{
		Status: "success",
		Data:   healthStatus,
	}

	c.JSON(http.StatusOK, response)
}
