package db

import (
	"context"
	"os"
	"time"

	"backend/core"
	"backend/settings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var dbClient *mongo.Client
var dbInstance *mongo.Database
var activityCollection *mongo.Collection
var userCollection *mongo.Collection

var (
	appSettings = settings.Get()
	log         = core.GetLogger()
)

func GetInstance() *mongo.Database {
	return dbInstance
}

func GetURI() string {
	uri := os.Getenv("MONGO_URL")
	return uri
}
func Connect() {
	uri := GetURI()
	// fmt.Println(uri)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	dbClient = client

	dbInstance = client.Database("backend")

	// println("[database] connected")
}

func _getCol(col string) *mongo.Collection {
	return dbInstance.Collection(col)
}

func _ctx() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}
