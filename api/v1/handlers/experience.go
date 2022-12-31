package handlers

import (
	"golang-rest-api-portfolio/api/infra/storage"
	"golang-rest-api-portfolio/api/v1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetExperiences(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, storage.GetExperiences())
}

func PostExperience(c *gin.Context) {
	var newExperience models.Experience

	if err := c.BindJSON(&newExperience); err != nil {
		return
	}

	storage.AddExperience(newExperience)
	c.IndentedJSON(http.StatusCreated, newExperience)
}

func GetExperienceByID(c *gin.Context) {
	id := c.Param("id")

	if experience := storage.GetExperienceByID(id); experience == nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "experience not found"})
	} else {
		c.IndentedJSON(http.StatusOK, experience)
	}
}

func DeleteExperience(c *gin.Context) {
	id := c.Param("id")

	if deleted_id := storage.RemoveExperience(id); deleted_id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "experience id not found"})
	} else {
		c.IndentedJSON(http.StatusOK, deleted_id)
	}
}
