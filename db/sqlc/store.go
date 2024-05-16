package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	TransferTx(ctx context.Context, arg CreateTransferParams) (TransferTxResult, error)

}
type SQLStore struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) Store {
	//remeber an interface takes the type of any struct that calls it. 
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store * SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	txQuery  := New(tx)
	txErr := fn(txQuery)
	if txErr != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("txErr: %v\nrbErr: %v", txErr, rbErr)
		}
		return txErr
	}
	return tx.Commit()
}

//CreateTxParams has the same field as CreateTransferParams so just use it

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// TransferTx performs a money transfer from one account to another
// It creates a transfer record, adds account entries, and updates
// the balance of both accounts involved within a single database transaction
func (store *SQLStore) TransferTx(ctx context.Context, arg CreateTransferParams) (TransferTxResult, error) {
	var result TransferTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var transTxErr error


		result.Transfer, transTxErr = q.CreateTransfer(ctx, arg)
		if transTxErr != nil {
			return transTxErr
		}


		result.FromEntry, transTxErr = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if transTxErr != nil {
			return transTxErr
		}

		result.ToEntry, transTxErr = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if transTxErr != nil {
			return transTxErr
		}
		//updating account
		//Update the Balance of the entryTo account

		if result.Transfer.FromAccountID > result.Transfer.ToAccountID {
			result.FromAccount, result.ToAccount, transTxErr = updateTopAccBalance(ctx, q, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
			if transTxErr != nil {
				return transTxErr
			}
		}
		if result.Transfer.FromAccountID < result.Transfer.ToAccountID {
			result.ToAccount, result.FromAccount, transTxErr = updateTopAccBalance(ctx, q, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
			if transTxErr != nil {
				return transTxErr
			}
		}
		return nil

	})
	return result, err
}
// updates the account with a higher account id
func updateTopAccBalance(ctx context.Context, q *Queries, account1, amount1,  account2, amount2 int64)(Account, Account, error){
	acc1Params := UpdateAccountBalanceParams{
		ID: account1,
		Amount: amount1,
	}
	acc2Params := UpdateAccountBalanceParams{
		ID: account2,
		Amount: amount2,
	}
	fromAccount, hErr := q.UpdateAccountBalance(ctx, acc1Params)
	toAccount, lErr := q.UpdateAccountBalance(ctx, acc2Params)
	if hErr != nil {
		return Account{}, Account{}, hErr
	}
	if lErr != nil {
		return Account{}, Account{}, lErr
	}
	return fromAccount, toAccount, nil
}