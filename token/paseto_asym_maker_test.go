package token

import (
	"crypto/rand"
	"github.com/fsobh/auth/util"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/ed25519"
	"testing"
	"time"
)

func TestNewAsymPasetoMaker(t *testing.T) {
	publicKey, privateKey, _ := ed25519.GenerateKey(rand.Reader)
	maker, err := NewAsymPasetoMaker(privateKey, publicKey)

	if err != nil || maker == nil {
		t.Errorf("failed to create PasetoMaker: error %v, maker %v", err, maker)
	}
}

func TestCreateToken(t *testing.T) {
	publicKey, privateKey, _ := ed25519.GenerateKey(rand.Reader)
	maker, _ := NewAsymPasetoMaker(privateKey, publicKey)

	username := util.RandomOwner()
	duration := time.Minute
	role := util.UserRole

	_, _, err := maker.CreateToken(username, role, duration)
	if err != nil {
		t.Errorf("failed to create token: %v", err)
	}
}

func TestCreateVerifyToken(t *testing.T) {
	publicKey, privateKey, _ := ed25519.GenerateKey(rand.Reader)
	maker, _ := NewAsymPasetoMaker(privateKey, publicKey)

	username := util.RandomOwner()
	duration := time.Minute
	role := util.UserRole

	token, payload, _ := maker.CreateToken(username, role, duration)
	payload, err := maker.VerifyToken(token)

	if err != nil || payload == nil || payload.Username != username {
		t.Errorf("failed to verify token: payload %v, error %v", payload, err)
	}
}
func TestInvalidToken(t *testing.T) {
	publicKey, privateKey, _ := ed25519.GenerateKey(rand.Reader)
	maker, _ := NewAsymPasetoMaker(privateKey, publicKey)

	// Test with invalid token
	token := "invalid.token.string"
	payload, err := maker.VerifyToken(token)

	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}

func TestExpiredToken(t *testing.T) {
	publicKey, privateKey, _ := ed25519.GenerateKey(rand.Reader)
	maker, _ := NewAsymPasetoMaker(privateKey, publicKey)
	username := util.RandomOwner()
	role := util.UserRole

	// Test with token created with a duration of negative 1 second, this token should be expired at the time of creation
	token, _, err := maker.CreateToken(username, role, -time.Second)

	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)

	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
