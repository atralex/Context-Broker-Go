package controllers

import (
	"Context-Broker/db"
	"Context-Broker/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddEntity(c *gin.Context) {
	var entity entities.Entity
	var res string = entities.ValidateEntity(entity, c)
	if res != "ok" {
		c.JSON(http.StatusBadRequest, gin.H{"status": res})
		return
	}

	var client, ctx, cancel = db.ConnectToMongo()
	db.AddEntity(client, ctx, entity)
	defer cancel()
	//TODO: Agregar la entidad a la base de datos
	// Si todo está correcto, haz algo con la entidad (por ejemplo, guardar en DB).
	c.JSON(http.StatusOK, gin.H{"status": "Entidad recibida correctamente", "entity": entity})
}

func All(c *gin.Context) {
	var client, ctx, cancel = db.ConnectToMongo()
	var entities, err = db.GetAllEntities(client, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error al obtener las entidades"})
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, gin.H{"entities": entities})
}
