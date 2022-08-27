package service

import (
	"encoding/hex"
	"fmt"
	"math/rand"

	"golang.org/x/crypto/scrypt"
)

func hash_options() (int, int, int, int) {
	N := 32768
	r := 8
	p := 1
	key_length := 32

	return N, r, p, key_length
}

func hashPass(pass string) (string, error) {
	salt := make([]byte, 32)
	_, err := rand.Read(salt)

	if err != nil {
		return "", err
	}

	N, r, p, key_length := hash_options()

	shash, err := scrypt.Key([]byte(pass), salt, N, r, p, key_length)

	if err != nil {
		return "", err
	}

	hassed_pass := fmt.Sprintf("%s.%s", hex.EncodeToString(shash), hex.EncodeToString(salt))

	return hassed_pass, nil
}
