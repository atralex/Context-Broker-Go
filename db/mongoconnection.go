package db

import (
	"Context-Broker/entities"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongo() (*mongo.Client, context.Context, context.CancelFunc) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb:context@broker//mongo:27017"))
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

func AddEntity(client *mongo.Client, ctx context.Context, entity entities.Entity) error {
	typesCollection := client.Database("Context-Broker").Collection("types")
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
		var properties map[string]entities.DataType
		for name, prop := range entity.Properties {
			properties[name] = prop.Type
		}
		// Construir y agregar el esquema a 'schemas'
		schema := bson.M{"type": entity.Type, "properties": properties}
		_, err = schemasCollection.InsertOne(ctx, schema)
		if err != nil {
			return err
		}
	} else if err != nil {
		// Otro error al buscar el documento
		return err
	}

	// Agregar la entidad a la colección correspondiente a su tipo
	entitiesCollection := client.Database("Context-Broker").Collection("entities")
	_, err = entitiesCollection.InsertOne(ctx, entity)
	return err
}

func GetEntity(client *mongo.Client, ctx context.Context, id string) (entities.Entity, error) {
	entitiesCollection := client.Database("Context-Broker").Collection("entities")
	var entity entities.Entity
	err := entitiesCollection.FindOne(ctx, bson.M{"id": id}).Decode(&entity)
	return entity, err
}

func GetEntitiesByType(client *mongo.Client, ctx context.Context, entityType string) ([]entities.Entity, error) {
	entitiesCollection := client.Database("Context-Broker").Collection("entities")
	cursor, err := entitiesCollection.Find(ctx, bson.M{"type": entityType})
	if err != nil {
		return nil, err
	}
	var entitiesList []entities.Entity
	err = cursor.All(ctx, &entitiesList)
	return entitiesList, err
}

func GetAllEntities(client *mongo.Client, ctx context.Context) ([]entities.Entity, error) {
	entitiesCollection := client.Database("Context-Broker").Collection("entities")
	cursor, err := entitiesCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var entitiesList []entities.Entity
	err = cursor.All(ctx, &entitiesList)
	return entitiesList, err
}
