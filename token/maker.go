package token

import "time"

type Maker interface {
	// NOTE: creates a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	// NOTE: check if the tokne is valid or not
	VerifyToken(token string) (*Payload, error)
}
