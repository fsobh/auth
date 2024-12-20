package db

import (
	"context"
	"fmt"
)

// execTx executes a function within a database transaction
// starts with lower case e (unexported) because we do not want external packages to call it directly.
// instead, we'll provide an exported function for each specific transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {

	tx, err := store.connPool.Begin(ctx) // &sql.TxOptions{} lets you set isolation level and read only property

	if err != nil {
		return err
	}

	q := New(tx) // this works cuz New accepts Db.tx interface
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			//if rollback fails
			return fmt.Errorf("tx err: %v, rb error : %v", err, rbErr) // output and return the original err + rollback err
		}
		return err // if rollback successful , return original tx error
	}
	return tx.Commit(ctx)
}
