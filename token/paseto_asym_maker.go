package token

import (
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/ed25519"
	"time"
)

type AsymPasetoMaker struct {
	paseto     *paseto.V2
	privateKey ed25519.PrivateKey
	publicKey  ed25519.PublicKey
}

func NewAsymPasetoMaker(privateKey ed25519.PrivateKey, publicKey ed25519.PublicKey) (Maker, error) {
	maker := &AsymPasetoMaker{
		paseto:     paseto.NewV2(),
		privateKey: privateKey,
		publicKey:  publicKey,
	}

	return maker, nil
}

func (maker *AsymPasetoMaker) CreateToken(username string, role string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, role, duration)

	if err != nil {
		return "", payload, err
	}

	token, err := maker.paseto.Sign(maker.privateKey, payload, nil)

	return token, payload, err
}

func (maker *AsymPasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := maker.paseto.Verify(token, maker.publicKey, payload, nil)

	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, ErrExpiredToken
	}

	return payload, nil
}
