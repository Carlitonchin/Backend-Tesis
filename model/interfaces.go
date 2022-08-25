package model

import "context"

type UserService interface {
	GetById(ctx context.Context, id uint) (*User, error)
}

type UserRepository interface {
	FindById(ctx context.Context, id uint) (*User, error)
}
