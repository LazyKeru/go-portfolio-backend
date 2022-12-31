package routers

import (
	"golang-rest-api-portfolio/api/v1/handlers"

	"github.com/gin-gonic/gin"
)

func SetupPing(router *gin.RouterGroup) {
	router.GET("/ping", handlers.Ping)

}
