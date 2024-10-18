package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)


const minSecretSizeKey = 32

// JWTMaker is a JSON web token maker
type JWTMaker struct {
	secretKey	string
}

// NewJWTMaker creates new JWTMaker
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretSizeKey {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretSizeKey)
	}

	return &JWTMaker{secretKey}, nil
}


// CreateToken creates a new token for username and duration
func(maker *JWTMaker)	CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(maker.secretKey))
}

// VerifyToken checks the given token is valid or not
func(maker *JWTMaker)	VerifyToken(token string) (*Payload, error) {
	// TODO
}
