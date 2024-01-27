package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/google/uuid"
)

// Different types of error returend by token
var (
	ErrorExpiredToken error = fmt.Errorf("token is expired")
	ErrorInvalidToken error = fmt.Errorf("token is invalid")
)

// Payload contains the payload data of the token
type Payload struct {
	ID       uuid.UUID `json:"id"`
	UserName string    `json:"user_name"`
	jwt.RegisteredClaims
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(userName string, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:       tokenId,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.RegisteredClaims.ExpiresAt.Time) {
		return ErrorExpiredToken
	}
	return nil
}
