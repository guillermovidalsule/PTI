package helper

import (
	"context"
	"fmt"
	"log"
	"os"
	"restful-api/database"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Email    string
	Username string
	Uid      string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "Users")
var robotCollection *mongo.Collection = database.OpenCollection(database.Client, "Robots")
var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, username string, uid string) (signedToken string, signedRefreshToken string, err error) {
	data := &SignedDetails{
		Email:    email,
		Username: username,
		Uid:      uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshData := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, data).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshData).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprint("the token is invalid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprint("token is expired")
		msg = err.Error()
		return
	}
	return claims, msg
}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
	// Crear un contexto con un tiempo de espera de 100 segundos
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	// Crear un objeto de actualización para los tokens y la hora de actualización
	var updateObj primitive.D

	// Agregar los nuevos tokens y la hora de actualización al objeto de actualización
	updateObj = append(updateObj, bson.E{"token", signedToken})
	updateObj = append(updateObj, bson.E{"refresh_token", signedRefreshToken})

	Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{"updated_at", Updated_at})

	// Configurar las opciones de actualización (usar upsert para insertar un nuevo documento si no existe)
	upsert := true
	filter := bson.M{"user_id": userId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	// Actualizar el documento en la colección de usuarios con los nuevos tokens y la hora de actualización
	_, err := userCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{"$set", updateObj},
		},
		&opt,
	)
	defer cancel()

	if err != nil {
		log.Panic(err)
		return
	}

	return
}

/*func UpdateRobotTokens(robotID, signedToken string) {
	// Crear un contexto con un tiempo de espera de 100 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	// Crear un objeto de actualización para los tokens y la hora de actualización
	var updateObj primitive.D

	// Agregar los nuevos tokens y la hora de actualización al objeto de actualización
	Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{"updated_at", Updated_at})

	updateObj = append(updateObj, bson.E{"token", signedToken})

	// Configurar las opciones de actualización (usar upsert para insertar un nuevo documento si no existe)
	upsert := true
	filter := bson.M{"robot_id": robotID}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	// Actualizar el documento en la colección de robots con el nuevo token y la hora de actualización
	_, err := robotCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{"$set", updateObj},
		},
		&opt,
	)
	defer cancel()
	if err != nil {
		log.Panic(err)
		return
	}

	return
}
*/
