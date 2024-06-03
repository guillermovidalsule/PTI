package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"restful-api/database"
	helper "restful-api/helpers"
	"restful-api/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "Users")
var validate = validator.New()

// HashPassword toma una contraseña como entrada y devuelve su hash utilizando el algoritmo bcrypt.
func HashPassword(password string) string {

	// Genera un hash bcrypt a partir de la contraseña dada. El segundo argumento es el costo del hash.
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

// VerifyPassword verifica si la contraseña proporcionada coincide con la contraseña almacenada en la base de datos utilizando el algoritmo bcrypt.
func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	// Compara el hash almacenado con la contraseña proporcionada.
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""
	if err != nil {
		msg = fmt.Sprintf("login or password incorrect")
		check = false
	}
	return check, msg
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Crear un contexto con un tiempo de espera de 100 segundos
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		// Crear una variable para almacenar los datos del usuario
		var user models.User

		// Leer y vincular los datos JSON de la solicitud al objeto de usuario
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validar los datos del usuario
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		// Verificar si el correo electrónico del usuario ya está en uso
		_, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking the email"})
			return
		}

		// Generar un hash de la contraseña del usuario
		password := HashPassword(*user.Password)
		user.Password = &password

		// Establecer la fecha y hora de creación y actualización del usuario
		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		// Generar un ID único para el usuario y asignarlo
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()

		// Generar tokens de acceso y actualización para el usuario
		token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.Username, *&user.User_id)
		user.Token = &token
		user.Refresh_token = &refreshToken

		// Insertar el usuario en la base de datos
		_, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		// Manejar el éxito y devolver el token de acceso
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Crear un contexto con un tiempo de espera de 100 segundos
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		// Declarar variables para almacenar el usuario de la solicitud y el usuario encontrado en la base de datos
		var user models.User
		var foundUser models.User

		// Leer y vincular los datos JSON de la solicitud al objeto de usuario
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Buscar al usuario en la base de datos por nombre de usuario
		err := userCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "username or password incorrect"})
			return
		}

		// Verificar si la contraseña proporcionada coincide con la contraseña almacenada
		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if passwordIsValid != true {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		// Generar nuevos tokens de acceso y actualización para el usuario
		token, refreshToken, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.Username, *&foundUser.User_id)
		// Actualizar los tokens de acceso y actualización en la base de datos
		helper.UpdateAllTokens(token, refreshToken, foundUser.User_id)

		// Devolver el usuario encontrado con un código de estado 200 OK
		c.JSON(http.StatusOK, foundUser)
	}

}

func Existe() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Crear un contexto con un tiempo de espera de 100 segundos
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		// Declarar variables para almacenar el usuario de la solicitud y el usuario encontrado en la base de datos
		var user models.User
		var foundUser models.User

		// Leer y vincular los datos JSON de la solicitud al objeto de usuario
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Buscar al usuario en la base de datos por nombre de usuario
		err := userCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "username or password incorrect"})
			return
		}

		// Verificar si la contraseña proporcionada coincide con la contraseña almacenada
		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if passwordIsValid != true {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusOK, foundUser)
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Crear un contexto con un tiempo de espera de 100 segundos
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		// Declarar variables para almacenar el usuario de la solicitud y el usuario encontrado en la base de datos
		var user models.User
		var foundUser models.User

		// Leer y vincular los datos JSON de la solicitud al objeto de usuario
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Buscar al usuario en la base de datos por nombre de mail
		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "username or password incorrect"})
			return
		}

		// Verificar si la contraseña proporcionada coincide con la contraseña almacenada
		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if passwordIsValid != true {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		//Elimina todos los robots de user
		_, err = robotCollection.DeleteMany(ctx, bson.M{"user_id": foundUser.User_id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user's robots"})
			return
		}

		_, err = userCollection.DeleteOne(ctx, bson.M{"user_id": foundUser.User_id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
	}
}
