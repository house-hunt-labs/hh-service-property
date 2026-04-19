package config

import (
    "context"
    "log"
    "os"

    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
    MongoClient *mongo.Client
    Database    *mongo.Database
}

func LoadConfig() (*Config, error) {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using system env")
    }

    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        mongoURI = "mongodb://localhost:27017"
    }

    dbName := os.Getenv("DB_NAME")
    if dbName == "" {
        dbName = "house_hunt"
    }

    clientOptions := options.Client().ApplyURI(mongoURI)
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        return nil, err
    }

    err = client.Ping(context.TODO(), nil)
    if err != nil {
        return nil, err
    }

    log.Println("Connected to MongoDB")

    db := client.Database(dbName)

    return &Config{
        MongoClient: client,
        Database:    db,
    }, nil
}