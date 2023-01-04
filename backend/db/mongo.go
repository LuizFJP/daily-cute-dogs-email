package db

import (
	"context"
	"daily-cute-dogs-email/backend/models"
	"errors"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type DB struct {
	client     *mongo.Client
	collection *mongo.Collection
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
	collection := client.Database(os.Getenv("MONGO_DB")).Collection(os.Getenv("MONGO_COLLECTION"))

	return &DB{client, collection}
}

func (d *DB) Finish() {
	if err := d.client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}
}

func (d *DB) GetEmails() ([]models.Subscriber, error) {
	cursor, err := d.collection.Find(context.Background(), bson.M{})
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
	checkEmail, err := d.checkEmailAlreadyExists(email)
	if err != nil {
		return err
	}
	if checkEmail {
		return errors.New("email já existe, adicione outro")
	} else {
		_, err = d.collection.InsertOne(context.TODO(), bson.M{"email": email})
		if err != nil {
			return err
		}
		return nil
	}
}

func (d *DB) DeleteEmail(email string) error {
	checkEmail, err := d.checkEmailAlreadyExists(email)
	if err != nil {
		return err
	}
	if !checkEmail {
		return errors.New("email não encontrado")
	} else {
		_, err = d.collection.DeleteOne(context.TODO(), bson.M{"email": email})
		if err != nil {
			return err
		}
		return nil
	}
}

func (d *DB) checkEmailAlreadyExists(email string) (bool, error) {
	cur, err := d.collection.Find(context.TODO(), bson.M{"email": email})
	if err != nil {
		return true, err
	}
	if cur.RemainingBatchLength() != 0 {
		return true, nil
	}
	return false, nil
}
