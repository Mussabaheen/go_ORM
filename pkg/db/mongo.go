package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoClient struct {
	db *mongo.Database
}

var Collection = "bookstore"

type Model any

type MongoClient interface {
	Ping()
	InsertOne(ctx context.Context, item Model) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, filter any, result Model) error
}

func NewMongoClient() *mongoClient {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Panic("unable to connect with mongo", err)
	}
	db := client.Database("books")
	return &mongoClient{
		db: db,
	}
}

func (mc *mongoClient) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := mc.db.Client().Ping(ctx, readpref.Primary())
	if err != nil {
		log.Panic("unable to connect to mongodb", err)
	}
}

func (mc *mongoClient) InsertOne(ctx context.Context, item Model) (*mongo.InsertOneResult, error) {
	return mc.db.Collection(Collection).InsertOne(ctx, item)
}

func (mc *mongoClient) FindOne(ctx context.Context, filter any, result Model) error {
	return mc.db.Collection(Collection).FindOne(ctx, filter).Decode(result)
}
