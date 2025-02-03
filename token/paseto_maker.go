package token

import (
	"fmt"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
	"time"
)

type PasetoMaker struct {
	paseto      *paseto.V2
	symmetricKy []byte
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {

	//Make sure the key is of same length as a paseto symmetric key
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size : must be %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:      paseto.NewV2(),
		symmetricKy: []byte(symmetricKey),
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(username string, role string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, role, duration)
	if err != nil {
		return "", payload, err
	}

	token, err := maker.paseto.Encrypt(maker.symmetricKy, payload, nil)

	return token, payload, err
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {

	//decrypt payload
	payload := &Payload{}
	err := maker.paseto.Decrypt(token, maker.symmetricKy, payload, nil)

	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()

	if err != nil {
		return nil, err
	}

	return payload, nil
}
