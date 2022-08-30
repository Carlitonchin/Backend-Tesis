package repository

import (
	"context"
	"fmt"
	"time"

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

func (s *redisTokenRepository) SetNewRefreshToken(ctx context.Context, userId string,
	tokenId string, expiresIn time.Duration) error {

	key := fmt.Sprintf("%s:%s", userId, tokenId)
	err := s.RedisClient.Set(ctx, key, 0, expiresIn).Err()

	return err
}

func (s *redisTokenRepository) DeleteRefreshToken(ctx context.Context,
	userId string, prevTokenId string) error {

	key := fmt.Sprintf("%s:%s", userId, prevTokenId)
	err := s.RedisClient.Del(ctx, key).Err()

	return err
}
