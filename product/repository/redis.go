package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"product/domain"
)

var (
	RedisClient *redis.Client
	RedisCtx    context.Context
)

func (r Repository) PublishMessageRedis(channel string, message []byte) *domain.Errors {
	res := RedisClient.Publish(RedisCtx, channel, message)
	_, err := res.Result()
	if err != nil {
		return domain.SetError(domain.CantPublishErr, err.Error())
	}
	return nil
}
