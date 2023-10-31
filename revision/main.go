package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"revision/domain"
	"revision/received/api"
	"revision/received/messaging"
)

type RedisRequirements struct {
	RedisClient *redis.Client
	RedisCtx    context.Context
}

var (
	Repo           domain.RepositoryInterface
	PORT           string
	Gin            *gin.Engine
	RestHandler    *api.RestHandler
	MessageHandler *messaging.MassageHandler
	RedisReq       *RedisRequirements
)

func main() {
	//subscribe
	go MessageHandler.SubscribeInsertRevisionMessageRedis()

	//define routes
	Gin.GET("/api/revision/products/id/:product-id/version-no/:version-no", RestHandler.GetRevisionByProductIDAndNo)
	Gin.GET("/api/revision/products/id/:product-id/revisions", RestHandler.GetAllRevisionsOfOneProduct)

	//run gin
	err := Gin.Run(PORT)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("-> Server running on ", PORT)
}
