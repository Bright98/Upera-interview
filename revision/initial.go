package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"revision/domain"
	"revision/received/api"
	"revision/received/messaging"
	"revision/repository"
	"strconv"
)

func getServerPort() string {
	return ":" + os.Getenv("PORT")
}
func mongoConnection() error {
	//get mongo requirements from env file
	timeout := os.Getenv("MONGO_TIMEOUT")
	mongoUrl := os.Getenv("MONGO_URL")
	database := os.Getenv("MONGO_DATABASE")
	mongoUsername := os.Getenv("DB_USERNAME")
	mongoPassword := os.Getenv("DB_PASSWORD")
	timeoutInt, err := strconv.Atoi(timeout)
	if err != nil {
		return err
	}

	//mongo connection
	return repository.MongoConnection(mongoUrl, database, mongoUsername, mongoPassword, timeoutInt)
}
func redisConnection() error {
	//get redis requirements from env file
	redisAddress := os.Getenv("REDIS_ADDRESS")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	db := os.Getenv("REDIS_DB")
	dbInt, err := strconv.Atoi(db)
	if err != nil {
		return err
	}

	//redis connection
	client, ctx, err := repository.RedisConnection(redisAddress, redisPassword, dbInt)
	if err != nil {
		return err
	}

	redisReq := &RedisRequirements{}
	redisReq.RedisClient = client
	redisReq.RedisCtx = ctx
	RedisReq = redisReq

	return nil
}

func init() {
	//load env file
	err := domain.LoadEnvFile()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("-> Environments loaded")

	//mongo connection
	err = mongoConnection()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("-> MongoDB connected")

	//redis connection
	err = redisConnection()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("-> Redis connected")

	//handle directory connection
	Repo = repository.NewRepository()
	service := domain.NewService(Repo)
	RestHandler = api.NewRestApi(service)
	MessageHandler = messaging.NewRedisMessage(RedisReq.RedisClient, RedisReq.RedisCtx, service)
	fmt.Println("-> Directory connection checked")

	PORT = getServerPort()
	Gin = gin.Default()
}
