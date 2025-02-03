package token

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size : must be atleast %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

// CreateToken Create a token for a specific username with a duration
func (maker *JWTMaker) CreateToken(username string, role string, duration time.Duration) (string, *Payload, error) {

	payload, err := NewPayload(username, role, duration)
	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return token, payload, err
}

// VerifyToken Check if the input token is valid or not
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {

	// a key function receives a parsed BUT unverified token.
	// Lets you use properties in key header (make sure the signing algorithm in it matches the algorithm you used)
	// This prevents the trivial mechanism exploit.

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		//try to convert the token's signing method to HMAC signing method since we used HS256, which is of SigningHmac type
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			//return error if conversion fails
			return nil, ErrInvalidToken
		}
		// return the maker's secret
		return []byte(maker.secretKey), nil
	}

	// parse the token with claims passing in the token, an empty payload, and a key function
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)

	if err != nil {
		// If an error occurs, we check if it's a validation error (by trying to convert it to one)
		verr, ok := err.(*jwt.ValidationError)

		//If it IS a validation error, then we check what kind it is:

		//if it's an ErrExpiredToken error (the Payloads Valid function returns an error )
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		//else, it can only be an invalid token error (keyFunc returned an error)
		return nil, ErrInvalidToken
	}

	//We attempt to get the payload data by converting the JWT token into a payload Object
	payload, ok := jwtToken.Claims.(*Payload)

	if !ok {
		//if the payload couldn't be converted successfully, return invalid token error
		return nil, ErrInvalidToken
	}
	//else, return the payload object and a nil error

	return payload, nil
}
