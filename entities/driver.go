package entities

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func connectToMongo() (*mongo.Client, context.Context, context.CancelFunc) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client, ctx, cancel
}

func addEntity(client *mongo.Client, ctx context.Context, entity Entity) error {
	typesCollection := client.Database("yourDatabaseName").Collection("types")
	schemasCollection := client.Database("yourDatabaseName").Collection("schemas")

	// Verificar si el tipo ya existe en 'types'
	var result interface{}
	err := typesCollection.FindOne(ctx, bson.M{"type": entity.Type}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// El tipo no existe, agregar a 'types'
		_, err := typesCollection.InsertOne(ctx, bson.M{"type": entity.Type})
		if err != nil {
			return err
		}

		// Construir y agregar el esquema a 'schemas'
		schema := bson.M{"type": entity.Type, "properties": entity.Properties}
		_, err = schemasCollection.InsertOne(ctx, schema)
		if err != nil {
			return err
		}
	} else if err != nil {
		// Otro error al buscar el documento
		return err
	}

	// Agregar la entidad a la colección correspondiente a su tipo
	entitiesCollection := client.Database("yourDatabaseName").Collection(entity.Type)
	_, err = entitiesCollection.InsertOne(ctx, entity)
	return err
}
