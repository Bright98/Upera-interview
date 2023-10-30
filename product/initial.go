package product

import (
	"fmt"
	"log"
	"os"
	"product/api"
	"product/domain"
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
	mongoUsername := os.Getenv("DB_USERNAME")
	mongoPassword := os.Getenv("DB_PASSWORD")
	timeoutInt, err := strconv.Atoi(timeout)
	if err != nil {
		return err
	}

	//mongo connection
	err = repository.MongoConnection(mongoUrl, database, mongoUsername, mongoPassword, timeoutInt)
	if err != nil {
		return err
	}

	return nil
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

	err = mongoConnection()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("-> MongoDB connected")

	PORT = getServerPort()
}
