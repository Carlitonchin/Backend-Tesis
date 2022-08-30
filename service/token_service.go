package service

import (
	"context"
	"crypto/rsa"
	"strconv"

	"github.com/Carlitonchin/Backend-Tesis/model"
)

type tokenService struct {
	TokenRepository      model.TokenRepository
	PrivateKey           *rsa.PrivateKey
	PublicKey            *rsa.PublicKey
	RefreshSecret        string
	IDExipirationSec     int64
	RefreshExpirationSec int64
}

type TSConfig struct {
	TokenRepository      model.TokenRepository
	PrivateKey           *rsa.PrivateKey
	PublicKey            *rsa.PublicKey
	RefreshSecret        string
	IDExpirationSec      int64
	RefreshExpirationSec int64
}

func NewTokenService(c *TSConfig) model.TokenService {
	return &tokenService{
		PrivateKey:           c.PrivateKey,
		PublicKey:            c.PublicKey,
		RefreshSecret:        c.RefreshSecret,
		TokenRepository:      c.TokenRepository,
		IDExipirationSec:     c.IDExpirationSec,
		RefreshExpirationSec: c.RefreshExpirationSec,
	}
}

func (s *tokenService) GetNewPairFromUser(
	ctx context.Context,
	user *model.User,
	prevTokenId string) (*model.TokenPair, error) {

	//function body starts here
	id_token, err := generateIDToken(user, s.PrivateKey, s.IDExipirationSec)
	if err != nil {
		return nil, err
	}

	refresh_token, err := generateRefreshToken(user.ID, s.RefreshSecret, s.RefreshExpirationSec)

	if err != nil {
		return nil, err
	}
	userId_str := strconv.FormatUint(uint64(user.ID), 10)
	err = s.TokenRepository.SetNewRefreshToken(ctx, userId_str, id_token, refresh_token.ExipresIn)

	if err != nil {
		return nil, err
	}

	if prevTokenId != "" {
		err = s.TokenRepository.DeleteRefreshToken(ctx, userId_str, prevTokenId)

		if err != nil {
			return nil, err
		}
	}

	return &model.TokenPair{
		IDToken:      id_token,
		RefreshToken: refresh_token.SS,
	}, nil
}
