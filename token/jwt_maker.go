package token

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

// JWTMaker is JSON Wed Token Maker
type JWTMaker struct {
	secretKey string
}

// NewJWTMaker creates a new JWTMaker
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be atleast %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

// CreateToke create a token for specific userName and duration
func (maker *JWTMaker) CreateToken(userName string, duration time.Duration) (string, error) {
	mySigningKey := []byte(maker.secretKey)
	payload, err := NewPayload(userName, duration)
	if err != nil {
		return "", err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString(mySigningKey)

}

// VerifyToken checks if token is valid or not
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrorInvalidToken
		}
		return []byte(maker.secretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			return nil, ErrorExpiredToken
		}
		return nil, ErrorInvalidToken
	}
	if !jwtToken.Valid {
		return nil, ErrorInvalidToken
	}
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrorInvalidToken
	}
	return payload, nil
}
