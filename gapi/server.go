package gapi

import (
	"fmt"
	db "github.com/fsobh/auth/db/sqlc"
	"github.com/fsobh/auth/pb"
	"github.com/fsobh/auth/util"
	"github.com/fsobh/auth/worker"
	"github.com/fsobh/token"
)

// Server define http server here 1.
type Server struct {
	pb.UnimplementedAuthServer // enable forward compatibility meaning : server can accept calls to RPC's before they are actually implemented
	config                     util.Config
	store                      db.Store
	tokenMaker                 token.Maker
	taskDistributor            worker.TaskDistributor
}

// NewServer Creates a new gRPC server
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {

	accessTokenMaker, err := token.NewPasetoV2Public(config.PasetoPrivateKeyStr, config.PasetoPublicKeyStr)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      accessTokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil

}
