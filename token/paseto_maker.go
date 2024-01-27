package token

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

// PasetoMaker is Paseto token maker
type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

// NewPasetoMaker creates a new PasetoMaker
func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be atleast %d characters", chacha20poly1305.KeySize)
	}
	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}
	return maker, nil
}

// CreateToke create a token for specific userName and duration
func (maker *PasetoMaker) CreateToken(userName string, duration time.Duration) (string, error) {
	payload, err := NewPayload(userName, duration)
	if err != nil {
		return "", nil
	}
	return maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
}

// VerifyToken checks if token is valid or not
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrorInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, ErrorExpiredToken
	}

	return payload, nil

}
