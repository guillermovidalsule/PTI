package controllers

import (
	"context"
	"fmt"
	"net/http"
	"restful-api/database"

	"restful-api/models"
	"time"

	"github.com/gin-gonic/gin"
	//"github.com/go-playground/validator/v10"

	"go.mongodb.org/mongo-driver/mongo"
	//"golang.org/x/crypto/bcrypt"
)

var testCollection *mongo.Collection = database.OpenCollection(database.Client, "Test")

func AddParam() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Crear un contexto con un tiempo de espera de 100 segundos
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		// Crear una variable para almacenar los datos del usuario
		var test models.Test

		// Leer y vincular los datos JSON de la solicitud al objeto de usuario
		if err := c.BindJSON(&test); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Insertar el usuario en la base de datos
		_, insertErr := testCollection.InsertOne(ctx, test)
		if insertErr != nil {
			msg := fmt.Sprintf("nada")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		// Manejar el éxito y devolver el token de acceso
		c.JSON(http.StatusOK, gin.H{"cpu_freq": test.Cpu_freq})

	}
}
