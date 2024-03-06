package handlers

import (
	"finder/api/models"
	logging "finder/api/monitoring"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ExploreHandler(c *gin.Context, logger *logging.Logger) {
	targetPath := c.Query("path")
	if targetPath == "" {
		logger.Error("Bad Request: The 'path' parameter is missing or empty.")

		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status: "error",
			Error: models.Error{
				Code:      http.StatusBadRequest,
				Message:   "Bad Request",
				Details:   "The 'path' parameter is missing or empty.",
				Timestamp: time.Now().UTC().Format(time.RFC3339),
			},
		})
		return
	}

	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		logger.Error("Not Found: The specified directory does not exist - " + targetPath)

		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Status: "error",
			Error: models.Error{
				Code:      http.StatusNotFound,
				Message:   "Not Found",
				Details:   "The specified directory does not exist.",
				Timestamp: time.Now().UTC().Format(time.RFC3339),
			},
		})
		return
	}

	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	pageNumber, err := strconv.Atoi(page)
	if err != nil || pageNumber < 1 {
		logger.Error("Bad Request: Invalid 'page' parameter.")
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status: "error",
			Error: models.Error{
				Code:      http.StatusBadRequest,
				Message:   "Bad Request",
				Details:   "Invalid 'page' parameter.",
				Timestamp: time.Now().UTC().Format(time.RFC3339),
			},
		})
		return
	}

	limitNumber, err := strconv.Atoi(limit)
	if err != nil || limitNumber < 1 {
		logger.Error("Bad Request: Invalid 'limit' parameter.")
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status: "error",
			Error: models.Error{
				Code:      http.StatusBadRequest,
				Message:   "Bad Request",
				Details:   "Invalid 'limit' parameter.",
				Timestamp: time.Now().UTC().Format(time.RFC3339),
			},
		})
		return
	}

	entries := []models.DirectoryEntry{}

	err = filepath.Walk(targetPath, func(entryPath string, entryInfo os.FileInfo, err error) error {
		if entryPath == targetPath {
			return nil
		}

		relativePath, _ := filepath.Rel(targetPath, entryPath)
		entry := models.DirectoryEntry{
			Name:     relativePath,
			FullPath: entryPath,
			IsDir:    entryInfo.IsDir(),
			Size:     entryInfo.Size(),
			Mode:     entryInfo.Mode().String(),
			ModTime:  entryInfo.ModTime().Format(time.RFC3339),
		}

		entries = append(entries, entry)

		return nil
	})

	if err != nil {
		logger.Error("Internal Server Error: Failed to list directory - " + err.Error())

		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status: "error",
			Error: models.Error{
				Code:      http.StatusInternalServerError,
				Message:   "Internal Server Error",
				Details:   "Failed to list directory: " + err.Error(),
				Timestamp: time.Now().UTC().Format(time.RFC3339),
			},
		})
		return
	}

	startIndex := (pageNumber - 1) * limitNumber
	endIndex := startIndex + limitNumber
	if startIndex < 0 {
		startIndex = 0
	}
	if endIndex > len(entries) {
		endIndex = len(entries)
	}
	paginatedEntries := entries[startIndex:endIndex]

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status: "success",
		Data:   gin.H{"entries": paginatedEntries},
	})
}