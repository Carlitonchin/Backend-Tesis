package repository

import (
	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/go-redis/redis/v8"
)

type redisTokenRepository struct {
	RedisClient *redis.Client
}

func NewTokenRepository(redis_client *redis.Client) model.TokenRepository {
	return &redisTokenRepository{
		RedisClient: redis_client,
	}
}
