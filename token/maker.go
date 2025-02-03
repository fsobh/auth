package token

import "time"

// Maker interface will be used to manage the token creation and verification
// We will be creating support for both JWT and PASETO tokens
type Maker interface {

	// CreateToken Create a token for a specific username with a duration
	CreateToken(username string, role string, duration time.Duration) (string, *Payload, error)

	// VerifyToken Check if the input token is valid or not
	VerifyToken(token string) (*Payload, error)
}
