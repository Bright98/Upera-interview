package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"product/domain"
	"time"
)

type Repository struct {
}

func NewRepository() domain.RepositoryInterface {
	repo := &Repository{}
	return repo
}

// mongo
func MongoConnection(mongoUrl, database string, timeout int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		return err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	MongoTimeout = time.Duration(timeout) * time.Second
	MongoDatabase = client.Database(database)
	MongoClient = client

	return nil
}

// redis
func RedisConnection(address, password string, db int) error {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return err
	}

	RedisClient = client
	RedisCtx = context.Background()

	return nil
}
