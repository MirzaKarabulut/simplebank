package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	// CreateToken creates a new token for username and duration
	CreateToken(username string, duration time.Time) (string, error)

	// VerifyToken checks the given token is valid or not
	VerifyToken(token string) (*Payload, error)
}