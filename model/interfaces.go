package model

import (
	"context"
	"time"
)

type UserService interface {
	GetById(ctx context.Context, id uint) (*User, error)
	SignUp(ctx context.Context, user *User) error
}

type UserRepository interface {
	FindById(ctx context.Context, id uint) (*User, error)
	Create(ctx context.Context, user *User) error
}

type TokenService interface {
	GetNewPairFromUser(ctx context.Context, user *User, prevTokenId string) (*TokenPair, error)
}

type TokenRepository interface {
	SetNewRefreshToken(ctx context.Context, userId string, tokenId string, expiresIn time.Duration) error
	DeleteRefreshToken(ctx context.Context, userId string, prevTokenId string) error
}
