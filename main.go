package main

import (
	"Context-Broker/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/entity", func(c *gin.Context) {
		controllers.AddEntity(c)
	})
	router.Run(":8080")
}
