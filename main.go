package main

import (
	"golang-rest-api-portfolio/api/v1/routers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routers.SetupPing(router.Group("/api"))
	routers.SetupExperience(router.Group("/api/v1"))

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
