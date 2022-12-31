package routers

import (
	"golang-rest-api-portfolio/api/v1/handlers"

	"github.com/gin-gonic/gin"
)

func SetupExperience(router *gin.RouterGroup) {
	router.GET("/experiences", handlers.GetExperiences)
	router.GET("/experiences/:id", handlers.GetExperienceByID)
	router.POST("/experiences", handlers.PostExperience)

}
