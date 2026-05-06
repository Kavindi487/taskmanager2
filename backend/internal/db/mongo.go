package db

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Client   *mongo.Client
var Database *mongo.Database

func Connect(uri string, dbName string) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal("MongoDB connection error:", err)
    }

    if err := client.Ping(ctx, nil); err != nil {
        log.Fatal("MongoDB ping failed:", err)
    }

    log.Println("Connected to MongoDB!")
    Client   = client
    Database = client.Database(dbName)
}