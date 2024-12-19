package gapi

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	db "github.com/fsobh/auth/db/sqlc"
	"github.com/fsobh/auth/util"
	"github.com/fsobh/auth/worker"
	"github.com/fsobh/token"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/ed25519"
	"google.golang.org/grpc/metadata"
	"testing"
	"time"
)

func newTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)

	require.NoError(t, err)

	config := util.Config{
		PasetoPublicKeyStr:  hex.EncodeToString(publicKey),
		PasetoPrivateKeyStr: hex.EncodeToString(privateKey),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)
	return server
}

func newContextWithBearerToken(t *testing.T, token token.Maker, username string, duration time.Duration) context.Context {

	accessToken, _, err := token.CreateToken(username, duration)

	require.NoError(t, err)

	bearerToken := fmt.Sprintf("%s %s", authorizationBearer, accessToken)

	md := metadata.MD{
		authorizationHeader: []string{
			bearerToken,
		},
	}
	return metadata.NewIncomingContext(context.Background(), md)

}
