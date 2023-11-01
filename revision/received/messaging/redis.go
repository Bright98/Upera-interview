package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"revision/domain"
)

type MassageHandler struct {
	RedisClient      *redis.Client
	RedisCtx         context.Context
	ServiceInterface domain.ServiceInterface
}

func NewRedisMessage(redisClient *redis.Client, redisCtx context.Context, serviceInterface domain.ServiceInterface) *MassageHandler {
	return &MassageHandler{
		RedisClient:      redisClient,
		RedisCtx:         redisCtx,
		ServiceInterface: serviceInterface,
	}
}

// handle redis subscription
func (m MassageHandler) SubscribeInsertRevisionMessageRedis() {
	pubSub := m.RedisClient.Subscribe(m.RedisCtx, "insert-revision")
	defer pubSub.Close()
	for {
		m.HandleInsertRevisionMessage(pubSub)
	}
}
func (m MassageHandler) HandleInsertRevisionMessage(pubSub *redis.PubSub) {
	msg, err := pubSub.ReceiveMessage(m.RedisCtx)
	if err != nil {
		fmt.Println("[message error - receive]: ", err.Error())
		return
	}

	revision := &domain.Revisions{}
	err = json.Unmarshal([]byte(msg.Payload), revision)
	if err != nil {
		fmt.Println("[message error - unmarshal]: ", err.Error())
		return
	}

	insertedID, resErr := m.ServiceInterface.InsertRevisionService(revision)
	if resErr != nil {
		fmt.Println("[message error - service error]: ", resErr)
		return
	}
	fmt.Println("[message done - inserted id]: ", insertedID)
}
