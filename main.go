package main

import (
	"context"
	"database/sql"
	"errors"
	db "github.com/fsobh/auth/db/sqlc"
	_ "github.com/fsobh/auth/doc/statik"
	"github.com/fsobh/auth/gapi"
	"github.com/fsobh/auth/pb"
	"github.com/fsobh/auth/util"
	"github.com/fsobh/auth/worker"
	"github.com/fsobh/mail"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // need this for postgres migrate
	_ "github.com/golang-migrate/migrate/v4/source/file"       // so file paths work with migrate
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hibiken/asynq"
	_ "github.com/lib/pq" // VERY IMPORTANT FOR DB TO WORK ON SERVER
	"github.com/rakyll/statik/fs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
)

func runGrpcServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) {

	server, err := gapi.NewServer(config, store, taskDistributor)
	if err != nil {
		log.Fatal().Msg("Cannot create server")
	}

	grpcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)

	grpcServer := grpc.NewServer(grpcLogger)
	pb.RegisterAuthServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GrpcServerAddress)

	if err != nil {
		log.Fatal().Msg("Cannot create listener")
	}

	log.Info().Msgf("Start gRPC server on %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Msg("Cannot start GRPC server")
	}
}

func runGatewayServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) {

	server, err := gapi.NewServer(config, store, taskDistributor)
	if err != nil {
		log.Fatal().Msg("Cannot create server")
	}

	grpcMux := runtime.NewServeMux()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterAuthHandlerServer(ctx, grpcMux, server)

	if err != nil {
		log.Fatal().Msg("Cannot register server with grpc mux")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	statikFS, err := fs.New()

	if err != nil {
		log.Fatal().Msg("Cannot create statik file system")
	}

	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFS))
	mux.Handle("/swagger/", swaggerHandler)

	listener, err := net.Listen("tcp", config.HttpServerAddress)

	if err != nil {
		log.Fatal().Msg("Cannot create listener")
	}

	log.Info().Msgf("Start HTTP server on %s", listener.Addr().String())

	handler := gapi.HTTPLogger(mux)

	err = http.Serve(listener, handler)
	if err != nil {
		log.Fatal().Msg("Cannot start HTTP Gateway server")
	}
}

func runMigrations(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Msg("Cannot create migration instance")

	}
	if err = migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal().Msg("Cannot run migration")
	}
	log.Info().Msg("DB Migration successful")
}

func runTaskProcessor(config util.Config, redisOpt asynq.RedisClientOpt, store db.Store) {

	mailer := mail.NewSESSender(config.SESMailRegion, config.SESFromEmail)

	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, store, mailer)
	log.Info().Msg("Starting task processor")
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot start task processor")
	}
	log.Info().Msg("Task processor started")
}

func main() {

	//load env file
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Msg("Cannot load env variables")
	}

	if config.Enviroment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal().Msg("Can not connect to db:")
	}

	runMigrations(config.MigrationUrl, config.DBSource)

	store := db.NewStore(conn)

	redisOpts := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}

	taskDistributor := worker.NewRedisTaskDistributor(redisOpts)

	go runTaskProcessor(config, redisOpts, store)

	go runGatewayServer(config, store, taskDistributor)

	runGrpcServer(config, store, taskDistributor)
}
