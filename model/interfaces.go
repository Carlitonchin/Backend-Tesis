package model

import "context"

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
