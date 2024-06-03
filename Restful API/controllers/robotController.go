package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"restful-api/database"
	"restful-api/structures"

	"restful-api/models"

	"time"

	"github.com/gin-gonic/gin"
	//"github.com/go-playground/validator/v10"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	//"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

var robotCollection *mongo.Collection = database.OpenCollection(database.Client, "Robots")

/* VERSION 1 descartada por haman
func AltaRobot() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Crear un contexto con un tiempo de espera de 100 segundos
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		// Crear una variable para almacenar los datos del robot
		var robot models.Robot

		// Leer y vincular los datos JSON de la solicitud al objeto de robot
		if err := c.BindJSON(&robot); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validar los datos del robot
		validationErr := validate2.Struct(robot)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		// Verificar si el usuario ya tiene un robot con el mismo nombre
		filter := bson.M{
			"robotname": robot.Robotname, // Nombre del robot a verificar
			"userid":    robot.User_id,   // ID del usuario al que pertenece el robot
		}

		count, err := robotCollection.CountDocuments(ctx, filter)
		defer cancel()

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking the robot name"})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "a robot with the same name already exists for this user"})
			return
		}

		// Establecer la fecha y hora de creación y actualización del robot
		robot.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		robot.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		// Generar un ID único para el usuario y asignarlo
		robot.ID = primitive.NewObjectID()
		robot.Robot_id = robot.ID.Hex()

		//Inicializar estado incial del robot
		robot.Robotidle = false
		robot.Robotstate = false

		// Insertar el robot en la base de datos
		_, insertErr := robotCollection.InsertOne(ctx, robot)
		if insertErr != nil {
			msg := fmt.Sprintf("Robot item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		// Manejar el éxito y devolver mensaje de creación
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Robot '%s' created successfully for user '%s'", robot.Robotname, robot.User_id)})
	}
}
*/

/*VERSION 2 con route /robot/alta envias por JSON body userowner (es el username) y el robotname
func CreateRobotForUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Crear un contexto con un tiempo de espera de 100 segundos
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		// Crear una variable para almacenar los datos del robot
		var robot models.Robot
		var user models.User

		// Leer y vincular los datos JSON de la solicitud al objeto de robot
		if err := c.BindJSON(&robot); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}


		// Validar los datos del robot
		validationErr := validate.Struct(robot)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		// Verificar si el usuario ya tiene un robot con el mismo nombre
		filter := bson.M{
			"robotname": robot.Robotname, // Nombre del robot a verificar
			"userowner": robot.Userowner, // Username del usuario al que pertenece el robot
		}

		count, err := robotCollection.CountDocuments(ctx, filter)
		defer cancel()

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking the robot name"})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "a robot with the same name already exists for this user"})
			return
		}

		var ctx2, cancel2 = context.WithTimeout(context.Background(), 100*time.Second)



		err = userCollection.FindOne(ctx2, bson.M{"username": robot.Userowner}).Decode(&user)
		defer cancel2()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "username or password incorrect"})
			return
		}


		// Crear una instancia de Robot asociada al usuario
		robot.ID = primitive.NewObjectID()
		robot.Robotidle = false
		robot.Robotstate = false
		robot.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		robot.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		robot.Robot_id = robot.ID.Hex()
		robot.Token = user.Token // Utiliza el mismo token que el usuario

		// Insertar el robot en la base de datos
		_, err = robotCollection.InsertOne(ctx, robot)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating robot"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Robot created successfully",
			"user":    user,
			"robot":   robot,
		})
	}
}
*/

/* VERSION 3 /robot/alta/:token_user y por JSON body solo envias robotname
func CreateRobotForUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el nombre de usuario de los parámetros de la ruta
		token := c.Param("token_user")

		// Buscar el usuario en la base de datos
		var user models.User
		err := userCollection.FindOne(context.Background(), bson.M{"token": token}).Decode(&user)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to find user"})
			return
		}

		// Crear un contexto con un tiempo de espera de 100 segundos
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		// Leer y vincular los datos JSON de la solicitud al objeto de robot
		var robot models.Robot
		if err := c.BindJSON(&robot); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid robot data"})
			return
		}

		// Validar los datos del robot
		validationErr := validate.Struct(robot)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		// Verificar si el usuario ya tiene un robot con el mismo nombre
		filter := bson.M{
			"robotname": robot.Robotname, // Nombre del robot a verificar
			"token":     token,           // Username del usuario al que pertenece el robot
		}

		count, err := robotCollection.CountDocuments(ctx, filter)
		defer cancel()

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking the robot name"})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "a robot with the same name already exists for this user"})
			return
		}

		// Asignar parametros al robot
		robot.ID = primitive.NewObjectID()
		robot.Robotidle = false
		robot.Robotstate = false
		robot.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		robot.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		robot.Robot_id = robot.ID.Hex()
		robot.Token = user.Token // Utiliza el mismo token que el usuario
		robot.Userowner = user.Username

		// Insertar el nuevo robot en la base de datos
		_, err = robotCollection.InsertOne(context.Background(), robot)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create robot"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "robot created successfully"})
	}
}*/

// VERSION 4 Creo que es la version buena, en el middleware cojo el token de la cabezera de la peticion
// y como el token tiene los parametros email,uid, username hasheados con la clave luis de env solo me hace
// falta especificar por el body JSON el nombre de la nueva instancia robot
func CreateRobotForUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Obtener el usuario del contexto a través del middleware de autenticación, ya que el contexto se hereda gracias al c.Next()
		uid, _ := c.Get("user_id")

		var user models.User
		err := userCollection.FindOne(context.Background(), bson.M{"user_id": uid}).Decode(&user)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to find user"})
			return
		}

		// Crear un contexto con un tiempo de espera de 100 segundos
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		// Leer y vincular los datos JSON de la solicitud al objeto de robot
		var robot models.Robot
		if err := c.BindJSON(&robot); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid robot data"})
			return
		}

		// Validar los datos del robot
		validationErr := validate.Struct(robot)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		// Verificar si el usuario ya tiene un robot con el mismo nombre
		filter := bson.M{
			"robotname": robot.Robotname, // Nombre del robot a verificar
			"user_id":   uid,             // Username del usuario al que pertenece el robot
		}

		count, err := robotCollection.CountDocuments(ctx, filter)
		defer cancel()

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking the robot name"})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "a robot with the same name already exists for this user"})
			return
		}

		/*var coordenadas structures.Coord
		coordenadas.Latitud = "41°22'31.O''N"
		coordenadas.Longitud = "2°10'28.5''E"
		var ruta []structures.Coord
		ruta = append(ruta, coordenadas)

		var info structures.Macros
		info.Cpu_freq = 6.98
		info.Temperature = 9.98
		info.Velocity = 456.0*/

		// Asignar parametros al robot

		robot.ID = primitive.NewObjectID()
		robot.Robotidle = false
		robot.Robotstate = false
		robot.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		robot.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		robot.Robot_id = robot.ID.Hex()
		robot.User_id = user.User_id
		//robot.Robot_info = info
		//robot.Ruta = ruta

		// Insertar el nuevo robot en la base de datos
		_, err = robotCollection.InsertOne(context.Background(), robot)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create robot"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "robot created successfully"})
	}
}

// Ruta que te permite eliminar un robot especificado en los paraetros :robotname de la url, y que realciona con
// su propietario mediante el jwt de la peticion que esta en la cabezera y procesa el middleware
func DeleteRobotForUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		robotName := c.Param("robotname")
		// Obtener el usuario del contexto a través del middleware de autenticación, ya que el contexto se hereda gracias al c.Next()
		uid, _ := c.Get("user_id")

		// Buscar al usuario por su nombre de usuario
		filter := bson.M{"user_id": uid}
		var user models.User
		err := userCollection.FindOne(context.Background(), filter).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
			return
		}

		// Buscar al robot por su nombre y perteneciente al usuario
		filter = bson.M{"user_id": uid, "robotname": robotName}
		var robot models.Robot
		err = robotCollection.FindOne(context.Background(), filter).Decode(&robot)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Robot not found for this user"})
			return
		}

		// Eliminar el robot de la base de datos
		_, err = robotCollection.DeleteOne(context.Background(), filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting robot"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Robot deleted successfully from user"})
	}
}

// Listar todos los robots de un user
func ListUserRobots() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el nombre de usuario del contexto a través del middleware de autenticación
		uid, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Definir un filtro para buscar los robots del usuario por su nombre de usuario
		filter := bson.M{"user_id": uid}

		// Realizar una consulta a la base de datos para encontrar todos los robots del usuario
		cursor, err := robotCollection.Find(context.Background(), filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user's robots"})
			return
		}
		defer cursor.Close(context.Background())

		var robots []models.Robot
		// Iterar sobre los resultados del cursor y enviar cada robot al cliente
		for cursor.Next(context.Background()) {
			var robot models.Robot
			if err := cursor.Decode(&robot); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode robot data"})
				return
			}
			robots = append(robots, robot)
		}

		if err := cursor.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cursor error"})
			return
		}
		c.JSON(http.StatusOK, robots)
	}
}

func ConsultUserRobot() gin.HandlerFunc {
	return func(c *gin.Context) {
		robotName := c.Param("robotname")
		// Obtener el usuario del contexto a través del middleware de autenticación, ya que el contexto se hereda gracias al c.Next()
		uid, _ := c.Get("user_id")

		// Buscar al usuario por su nombre de usuario
		filter := bson.M{"user_id": uid}
		var user models.User
		err := userCollection.FindOne(context.Background(), filter).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
			return
		}

		// Buscar al robot por su nombre y perteneciente al usuario
		filter = bson.M{"user_id": uid, "robotname": robotName}
		var robot models.Robot
		err = robotCollection.FindOne(context.Background(), filter).Decode(&robot)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Robot not found for this user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"robot": robot})
	}

}

func RefreshMacrosRobot() gin.HandlerFunc {
	return func(c *gin.Context) {
		robotName := c.Param("robotname")
		// Obtener el usuario del contexto a través del middleware de autenticación, ya que el contexto se hereda gracias al c.Next()
		uid, _ := c.Get("user_id")

		// Crear un contexto con un tiempo de espera de 100 segundos
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Leer y vincular los datos JSON de la solicitud al objeto de robot
		var updatedRobotInfo structures.Macros
		if err := c.BindJSON(&updatedRobotInfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid robot info data"})
			return
		}
		fmt.Println("Updated Robot Info:", updatedRobotInfo)

		/*// Buscar al robot por su nombre y perteneciente al usuario
		filter := bson.M{"user_id": uid, "robotname": robotName}
		var robot models.Robot
		err := robotCollection.FindOne(ctx, filter).Decode(&robot)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Robot not found for this user"})
			return
		}*/

		// Actualizar la robot_info del robot en la base de datos
		filter := bson.M{"user_id": uid, "robotname": robotName}
		update := bson.M{"$set": bson.M{"robot_info": updatedRobotInfo}}
		_, err := robotCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update robot info"})
			return
		}

		var robot models.Robot
		err = robotCollection.FindOne(ctx, filter).Decode(&robot)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Robot not found for this user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"robot": robot})
	}

}
