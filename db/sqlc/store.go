package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

// SQLStore provides all functions to execute db queries and transactions
// For individual queries, we already have the Queries struct in db.go (generated by sqlc)
// each query does 1 operation per table so it doesnt support transactions (COMMIT/ROLLBACK)
// This file will extend its functionality to support Transactions
// by embedding it in the store struct (This is called composition)
// It is the preferred way to extend struct functionality in goLang instead of inheritance
type SQLStore struct {
	*Queries // all individual Query functions provided by `Queries` is available to Store struct
	connPool *pgxpool.Pool
}

// Store was created to be able to mock the DB for our tests.
// note : we updated sqlc.yaml with `emit_interface: true`, then we ran `make sqlc`
// so we won't have to copy everything that Queries implements into store interface (Check querier.go)
type Store interface {
	Querier
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
	VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error)
}

func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool), // New was generated by sqlc (creates and returns Queries obj)
	}
}
