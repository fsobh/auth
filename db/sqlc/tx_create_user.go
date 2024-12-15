package db

import "context"

type CreateUserTxParams struct {
	CreateUserParams
	//This is a callback function that will only be called after the user is created
	AfterCreate func(user User) error // output will determine weather to complete or rollback db tx
}

// CreateUserTxResult is the result of the transfer transaction
type CreateUserTxResult struct {
	User User
}

func (store *SQLStore) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	var result CreateUserTxResult // create empty result

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.User, err = q.CreateUser(ctx, arg.CreateUserParams)
		if err != nil {
			return err
		}
		//if user was successfully created, we call the callback function and pass in the arg it wants
		return arg.AfterCreate(result.User)

	})

	return result, err
}
