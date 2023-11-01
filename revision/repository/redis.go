package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"revision/domain"
	"strconv"
)

var (
	RedisClient *redis.Client
	RedisCtx    context.Context
)

func (r Repository) GetLastRevisionNoRedis(productID string) (int, *domain.Errors) {
	res := RedisClient.Get(RedisCtx, productID)
	value, err := res.Result()
	if err != nil {
		return 0, domain.SetError(domain.NotFoundErr, err.Error())
	}

	if value != "" {
		revisionNo, err := strconv.Atoi(value)
		if err != nil {
			return 0, domain.SetError(domain.InvalidationErr, err.Error())
		}
		return revisionNo, nil
	}
	return 0, nil
}
func (r Repository) SetLastRevisionNoRedis(productID string, revisionNo int) *domain.Errors {
	res := RedisClient.Set(RedisCtx, productID, revisionNo, redis.KeepTTL)
	_, err := res.Result()
	if err != nil {
		return domain.SetError(domain.NotFoundErr, err.Error())
	}
	return nil
}
