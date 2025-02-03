package token

import (
	"github.com/fsobh/auth/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestPasetoMaker(t *testing.T) {

	//create a new maker, PASSING IN A RANDOM KEY
	maker, err := NewPasetoMaker(util.RandomString(32))

	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute
	role := util.UserRole

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, payload, err := maker.CreateToken(username, role, duration)

	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.Equal(t, t, role, payload.Role)
	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)

}
func TestExpiredPasetoToken(t *testing.T) {

	//create a new maker, PASSING IN A RANDOM KEY
	maker, err := NewPasetoMaker(util.RandomString(32))

	require.NoError(t, err)

	//create a token, passing in a negative duration
	token, payload, err := maker.CreateToken(util.RandomOwner(), util.UserRole, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	//check if the token is valid and if we can decode the token into a payload
	payload, err = maker.VerifyToken(token)
	//Ensure the verification throws an error
	require.Error(t, err)
	//eEnsure the error thrown is of Expired token
	require.EqualError(t, err, ErrExpiredToken.Error())
	//Ensure the payload returned was nil
	require.Nil(t, payload)

}
