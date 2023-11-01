package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"product/domain"
	"product/received/api"
	"product/repository"
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
	timeoutInt, err := strconv.Atoi(timeout)
	if err != nil {
		return err
	}

	//mongo connection
	return repository.MongoConnection(mongoUrl, database, timeoutInt)
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
	return repository.RedisConnection(redisAddress, redisPassword, dbInt)
}

func init() {
	//load env file
	err := domain.LoadEnvFile()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("-> Environments loaded")

	//handle directory connection
	repo := repository.NewRepository()
	service := domain.NewService(repo)
	RestHandler = api.NewRestApi(service)
	fmt.Println("-> Directory connection checked")

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

	PORT = getServerPort()
	Gin = gin.Default()
}
