package service

import (
	"crypto/rsa"
	"math/rand"
	"time"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/dgrijalva/jwt-go"
)

//-------------------- ID TOKEN ------------------------------

type IDTokenClaims struct {
	User *model.User `json:"user"`
	jwt.StandardClaims
}

func generateIDToken(user *model.User, key *rsa.PrivateKey) (string, error) {
	now := time.Now().Unix()
	expired_time := now + 15*60 // 15 minutes

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

func generateRefreshToken(user_id uint, key string) (*RefreshToken, error) {
	currentTime := time.Now()
	tokenExp := currentTime.AddDate(0, 0, 3)

	tokenId := uint(rand.Uint64())

	claims := &RefreshTokenClaims{
		ID: user_id,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  currentTime.Unix(),
			ExpiresAt: tokenExp.Unix(),
			Id:        string(tokenId),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(key))

	if err != nil {
		return nil, err
	}

	return &RefreshToken{
		ID:        string(tokenId),
		SS:        ss,
		ExipresIn: tokenExp.Sub(currentTime),
	}, nil
}
