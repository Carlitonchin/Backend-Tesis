package service

import (
	"context"
	"crypto/rsa"

	"github.com/Carlitonchin/Backend-Tesis/model"
)

type TokenService struct {
	PrivateKey    *rsa.PrivateKey
	PublicKey     *rsa.PublicKey
	RefreshSecret string
}

type TSConfig struct {
	PrivateKey    *rsa.PrivateKey
	PublicKey     *rsa.PublicKey
	RefreshSecret string
}

func NewTokenService(c *TSConfig) model.TokenService {
	return &TokenService{
		PrivateKey:    c.PrivateKey,
		PublicKey:     c.PublicKey,
		RefreshSecret: c.RefreshSecret,
	}
}

func (s *TokenService) GetNewPairFromUser(
	ctx context.Context,
	user *model.User,
	prevTokenId string) (model.TokenPair, error) {

	//function body starts here
	panic("Not implemented yet")
}
