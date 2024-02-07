package main

import (
	"Context-Broker/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.POST("/entity", func(c *gin.Context) {
		var entity entities.Entity
		if err := c.ShouldBindJSON(&entity); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Print(entity.Properties)
		for _, prop := range entity.Properties {
			fmt.Print(prop.Type)
			validationFunc, ok := entities.Validators[string(prop.Type)]
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid property %s", prop.Type)})
				return
			}
			check := validationFunc(prop.Value)
			if !check {
				c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid property %s", prop.Type)})
				return
			}
		}
		//TODO: Agregar la entidad a la base de datos
		// Si todo está correcto, haz algo con la entidad (por ejemplo, guardar en DB).
		c.JSON(http.StatusOK, gin.H{"status": "Entidad recibida correctamente"})
	})
	router.Run(":8080")
}
