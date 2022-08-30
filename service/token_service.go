package service

import (
	"context"
	"crypto/rsa"

	"github.com/Carlitonchin/Backend-Tesis/model"
)

type tokenService struct {
	TokenRepository model.TokenRepository
	PrivateKey      *rsa.PrivateKey
	PublicKey       *rsa.PublicKey
	RefreshSecret   string
}

type TSConfig struct {
	TokenRepository model.TokenRepository
	PrivateKey      *rsa.PrivateKey
	PublicKey       *rsa.PublicKey
	RefreshSecret   string
}

func NewTokenService(c *TSConfig) model.TokenService {
	return &tokenService{
		PrivateKey:      c.PrivateKey,
		PublicKey:       c.PublicKey,
		RefreshSecret:   c.RefreshSecret,
		TokenRepository: c.TokenRepository,
	}
}

func (s *tokenService) GetNewPairFromUser(
	ctx context.Context,
	user *model.User,
	prevTokenId string) (*model.TokenPair, error) {

	//function body starts here
	id_token, err := generateIDToken(user, s.PrivateKey)
	if err != nil {
		return nil, err
	}

	refresh_token, err := generateRefreshToken(user.ID, s.RefreshSecret)

	if err != nil {
		return nil, err
	}

	//TODO: store refresh token

	return &model.TokenPair{
		IDToken:      id_token,
		RefreshToken: refresh_token.SS,
	}, nil
}
