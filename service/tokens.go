package service

import (
	"crypto/rsa"
	"time"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

//-------------------- ID TOKEN ------------------------------

type IDTokenClaims struct {
	User *model.User `json:"user"`
	jwt.StandardClaims
}

func generateIDToken(user *model.User, key rsa.PrivateKey) (string, error) {
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
	UID uuid.UUID `json:"uid"`
	jwt.StandardClaims
}

func generateRefreshToken(uid uuid.UUID, key string) (*RefreshToken, error) {
	currentTime := time.Now()
	tokenExp := currentTime.AddDate(0, 0, 3)

	tokenId, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	claims := &RefreshTokenClaims{
		UID: tokenId,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  currentTime.Unix(),
			ExpiresAt: tokenExp.Unix(),
			Id:        tokenId.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(key))

	if err != nil {
		return nil, err
	}

	return &RefreshToken{
		ID:        tokenId.String(),
		SS:        ss,
		ExipresIn: tokenExp.Sub(currentTime),
	}, nil
}
