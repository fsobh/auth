package token

import (
	"github.com/fsobh/auth/util"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestJWTMaker(t *testing.T) {

	//create a new maker, PASSING IN A RANDOM KEY
	maker, err := NewJWTMaker(util.RandomString(32))

	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute
	role := util.UserRole

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, payload, err := maker.CreateToken(username, role, duration)

	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.Equal(t, role, payload.Role)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)

}
func TestExpiredJWTToken(t *testing.T) {

	//create a new maker, PASSING IN A RANDOM KEY
	maker, err := NewJWTMaker(util.RandomString(32))

	require.NoError(t, err)

	//create a token, passing in a negative duration
	token, payload, err := maker.CreateToken(util.RandomOwner(), util.UserRole, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	//check if the token is valid and if we can decode the token into a payload
	payload, err = maker.VerifyToken(token)
	//Ensure the verification throws an error
	require.Error(t, err)
	//eEnsure the error thrown is of Expired token
	require.EqualError(t, err, ErrExpiredToken.Error())
	//Ensure the payload returned was nil
	require.Nil(t, payload)

}
func TestInvalidJWTTokenAlgNone(t *testing.T) {
	// Create a new payload
	payload, err := NewPayload(util.RandomOwner(), util.UserRole, time.Minute)
	require.NoError(t, err)

	//Generate a new token using the non signing method and pass in the payload
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	//we sign the token using the signed string method (we only pass in this param to allow it to sign using the none method)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)

}
