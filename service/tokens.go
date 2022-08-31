package service

import (
	"crypto/rsa"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/dgrijalva/jwt-go"
)

//-------------------- ID TOKEN ------------------------------

type IDTokenClaims struct {
	User *model.User `json:"user"`
	jwt.StandardClaims
}

func generateIDToken(user *model.User, key *rsa.PrivateKey, expiredIn int64) (string, error) {
	now := time.Now().Unix()
	expired_time := now + expiredIn

	claims := &IDTokenClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now,
			ExpiresAt: expired_time,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	ss, err := token.SignedString(key)

	return ss, err
}

//-------------------END IDTOKEN ----------------------

// --------------- Refresh token --------------------------
type RefreshToken struct {
	ID        string
	SS        string
	ExipresIn time.Duration
}

type RefreshTokenClaims struct {
	ID uint
	jwt.StandardClaims
}

func generateRefreshToken(user_id uint, key string, expiresIn int64) (*RefreshToken, error) {
	currentTime := time.Now()
	tokenExp := currentTime.Add(time.Duration(expiresIn) * time.Second)

	tokenId := uint(rand.Uint64())

	claims := &RefreshTokenClaims{
		ID: user_id,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  currentTime.Unix(),
			ExpiresAt: tokenExp.Unix(),
			Id:        strconv.FormatUint(uint64(tokenId), 10),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(key))

	if err != nil {
		return nil, err
	}

	return &RefreshToken{
		ID:        strconv.FormatUint(uint64(tokenId), 10),
		SS:        ss,
		ExipresIn: tokenExp.Sub(currentTime),
	}, nil
}

func validateIdToken(tokenString string, key *rsa.PublicKey) (*IDTokenClaims, error) {
	claims := &IDTokenClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Token invalido")
	}

	claims, ok := token.Claims.(*IDTokenClaims)

	if !ok {
		return nil, fmt.Errorf("No se pudo parsear los claims")
	}

	return claims, nil
}
