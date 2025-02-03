package db

import (
	"context"
	"github.com/fsobh/auth/util"
	"github.com/jackc/pgx/v5/pgxpool"

	"log"
	"os"
	"testing"
)

// Global variables in testing
var testStore Store

func TestMain(m *testing.M) {

	//load env file
	config, err := util.LoadConfig("../..") // pass in the path relative to this file

	if err != nil {
		log.Fatal("Cannot load env variables", err)
	}

	conPool, err := pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		log.Fatal("Can not connect to db:", err)
	}

	testStore = NewStore(conPool)

	os.Exit(m.Run())
}
