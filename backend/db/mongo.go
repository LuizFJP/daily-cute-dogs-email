package db

import (
	"context"
	"daily-cute-dogs-email/backend/models"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type DB struct {
	*mongo.Client
}

func Start() *DB {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(os.Getenv("MONGO_CREDENTIALS")).
		SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return &DB{client}
}

func (d *DB) Finish() {
	if err := d.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}
}

func (d *DB) GetEmails() ([]models.Subscriber, error) {
	collection := d.Database(os.Getenv("MONGO_DB")).Collection(os.Getenv("MONGO_COLLECTION"))
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return []models.Subscriber{}, err
	}
	defer cursor.Close(context.Background())

	var subscribers []models.Subscriber

	for cursor.Next(context.Background()) {
		var subscriber models.Subscriber
		err := cursor.Decode(&subscriber)
		if err != nil {
			return []models.Subscriber{}, err
		}
		subscribers = append(subscribers, subscriber)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return subscribers, nil
}

func (d *DB) AddEmail(email string) error {
	collection := d.Database(os.Getenv("MONGO_DB")).Collection(os.Getenv("MONGO_COLLECTION"))
	if _, err := collection.InsertOne(context.TODO(), bson.M{"email": email}); err != nil {
		return err
	}
	return nil
}