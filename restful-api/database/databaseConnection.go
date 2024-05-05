package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBInstance crea y devuelve una instancia de cliente de MongoDB
func DBinstance() *mongo.Client {

	// Cargar las variables de entorno desde el archivo .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Obtener la URL de conexión a MongoDB desde las variables de entorno
	MongoDb := os.Getenv("MONGODB_URL")

	// Crear un nuevo cliente de MongoDB usando la URL de conexión
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	// Contexto con tiempo de espera para la conexión
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// Conectar al servidor de MongoDB
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conected to MongoDB!")
	// Devolver la instancia del cliente de MongoDB
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	// Acceder a la base de datos con el nombre "Talos" utilizando el cliente de MongoDB
	// y obtener la colección con el nombre especificado por collectionName.
	var collection *mongo.Collection = client.Database("Talos").Collection(collectionName)
	// Devolver un puntero a la colección obtenida.
	// Esto permitirá a otras partes del programa acceder y manipular la colección.
	return collection
}
