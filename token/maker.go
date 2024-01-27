package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	// CreateToke create a token for specific userName and duration
	CreateToken(userName string, duration time.Duration) (string, error)

	// VerifyToken checks if token is valid or not
	VerifyToken(token string) (*Payload, error)
}
