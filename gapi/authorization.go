package gapi

import (
	"context"
	"fmt"
	"github.com/fsobh/auth/token"
	"google.golang.org/grpc/metadata"
	"strings"
)

const (
	authorizationHeader = "authorization"
	authorizationBearer = "bearer"
)

func (server *Server) authorizeUser(ctx context.Context) (*token.Payload, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("metadata is not provided")
	}
	values := md.Get(authorizationHeader)

	if len(values) == 0 {
		return nil, fmt.Errorf("authorization token is not provided")
	}

	authHeader := values[0]
	//<authorization-type> <authorization-data>
	fields := strings.Fields(authHeader)

	if len(fields) < 2 {
		return nil, fmt.Errorf("invalid authorization header")
	}

	authType := strings.ToLower(fields[0])

	if authType != authorizationBearer {
		return nil, fmt.Errorf("unsupported authorization type")
	}

	access_token := fields[1]

	payload, err := server.tokenMaker.VerifyToken(access_token)

	if err != nil {
		return nil, fmt.Errorf("access token is not valid")
	}

	return payload, nil

}
