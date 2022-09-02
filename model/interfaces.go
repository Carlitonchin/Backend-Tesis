package model

import (
	"context"
	"time"
)

type UserService interface {
	GetById(ctx context.Context, id uint) (*User, error)
	SignUp(ctx context.Context, user *User) error
	SignIn(ctx context.Context, user *User) error
	GetAllUsers(ctx context.Context) ([]User, error)
}

type UserRepository interface {
	FindById(ctx context.Context, id uint) (*User, error)
	Create(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	GetRoleById(ctx context.Context, id uint) (*Role, error)
	GetAllUsers(ctx context.Context) ([]User, error)
}

type TokenService interface {
	GetNewPairFromUser(ctx context.Context, user *User, prevTokenId string) (*TokenPair, error)
	ValidateIdToken(tokenString string) (*User, error)
	ValidateRefreshToken(refresh_token string) (*RefreshToken, error)
	SignOut(ctx context.Context, user_id uint) error
}

type TokenRepository interface {
	SetNewRefreshToken(ctx context.Context, userId string, tokenId string, expiresIn time.Duration) error
	DeleteRefreshToken(ctx context.Context, userId string, prevTokenId string) error
	DeleteUserRefreshTokens(ctx context.Context, userId string) error
}

type RoleService interface {
	GetRoles(ctx context.Context) ([]Role, error)
	GetRoleByName(ctx context.Context, role_name string) (*Role, error)
}

type RoleRepository interface {
	GetRoles(ctx context.Context) ([]Role, error)
	GetRoleByName(ctx context.Context, role_name string) (*Role, error)
}
