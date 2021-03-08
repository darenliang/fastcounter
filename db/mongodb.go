package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"time"
)

var Database *mongo.Database
var Context context.Context

func init() {
	uri := os.Getenv("MONGODB_URI")
	Context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(Context, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(Context, readpref.Primary()); err != nil {
		panic(err)
	}

	Database = client.Database("fastcounter")
}
