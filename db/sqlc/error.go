package db

import (
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	UniqueViolation     = "23505"
	ForeignKeyViolation = "23503"
)

var ErrRecordNotFound = pgx.ErrNoRows // doing this because sql.ErrNoWRows returns a different string message than pgx does, so we made this for our unit tests

var ErrUniqueViolation = &pgconn.PgError{
	Code: UniqueViolation,
}

func ErrorCode(err error) string {

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code
	}
	return ""
}
