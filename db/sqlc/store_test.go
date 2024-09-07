package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	accountTo := createRandomAccount(t)
	accountFrom := createRandomAccount(t)

	//number of tx
	n := 5
	amount := int64(10)
	txArgs := CreateTransferParams{
		FromAccountID: accountFrom.ID,
		ToAccountID:   accountTo.ID,
		Amount:        amount,
	}
	resultChan := make(chan TransferTxResult)
	errorChan := make(chan error)
	//create n transactions between two accounts
	for i := 0; i < n; i++ {
		//to monitor which transaction is blocking
		txName := fmt.Sprintf("tx %d", i + 1)
		go func() {
			ctx := context.WithValue(context.Background(), txKey, txName)
			txResult, err := testStore.TransferTx(ctx, txArgs)

			resultChan <- txResult
			errorChan <- err
		}()
	}

	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		txRes := <-resultChan
		txErr := <-errorChan

		require.NotEmpty(t, txRes)
		require.NoError(t, txErr)

		//check for transfers
		transfer := txRes.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, txArgs.Amount, transfer.Amount)
		require.Equal(t, txArgs.FromAccountID, transfer.FromAccountID)
		require.Equal(t, txArgs.ToAccountID, transfer.ToAccountID)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)
		

		_, err := testQueries.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		//check for entry
		fromEntry := txRes.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, accountFrom.ID, fromEntry.AccountID)
 

		toEntry := txRes.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, accountTo.ID, toEntry.AccountID)

		require.Equal(t, toEntry.Amount, -fromEntry.Amount)

		//check for account
		fromAccount := txRes.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, accountFrom.ID, fromAccount.ID)

		toAccount := txRes.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, accountTo.ID, toAccount.ID)


		//check th  e balance
		amountFromAccount := accountFrom.Balance - fromAccount.Balance
		amountToAccount := toAccount.Balance - accountTo.Balance
		require.Equal(t, amountFromAccount, amountToAccount)
		require.True(t, amountToAccount > 0)
		require.True(t, amountFromAccount%amount == 0)

		k := int(amountFromAccount / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	//check final balance and accounts exist
	updatedFromAccount, frmErr := testQueries.GetAccount(context.Background(), accountFrom.ID)
	require.NoError(t, frmErr)
	updatedToAccount, toErr := testQueries.GetAccount(context.Background(), accountTo.ID)
	require.NoError(t, toErr)
	require.NotEmpty(t, updatedFromAccount)
	require.NotEmpty(t, updatedToAccount)

	require.Equal(t, (accountFrom.Balance - int64(n)*amount), updatedFromAccount.Balance)
	require.Equal(t, (accountTo.Balance + int64(n)*amount), updatedToAccount.Balance)


}

func TestTransferTxDeadLock(t *testing.T) {
	accountTo := createRandomAccount(t)
	accountFrom := createRandomAccount(t)

	//number of tx
	n := 10
	amount := int64(10)
	txArgs := CreateTransferParams{
		FromAccountID: accountFrom.ID,
		ToAccountID:   accountTo.ID,
		Amount:        amount,
	} 
	
	errorChan := make(chan error)
	resChan := make(chan TransferTxResult)
	//create n transactions between two accounts
	for i := 0; i < n; i++ {
		if i % 2 != 1{
			txArgs.FromAccountID = accountTo.ID
			txArgs.ToAccountID = accountFrom.ID
		} else if i % 2 == 1 {
			txArgs.FromAccountID = accountFrom.ID
			txArgs.ToAccountID = accountTo.ID
		}
		go func(args CreateTransferParams) {
			resTx, err := testStore.TransferTx(context.Background(), args)
			errorChan <- err
			resChan <- resTx
		}(txArgs)
	}

	for i := 0; i < n; i++ {
		txErr := <- errorChan
		txRes := <- resChan
		require.NoError(t, txErr)
		require.NotEmpty(t, txRes)
	}
	//check final balance and accounts exist
	updatedFromAccount, frmErr := testQueries.GetAccount(context.Background(), accountFrom.ID)
	updatedToAccount, toErr := testQueries.GetAccount(context.Background(), accountTo.ID)
	require.NoError(t, frmErr)
	require.NoError(t, toErr)
	require.NotEmpty(t, updatedFromAccount)
	require.NotEmpty(t, updatedToAccount)

	require.Equal(t, (accountFrom.Balance), updatedFromAccount.Balance)
	require.Equal(t, (accountTo.Balance), updatedToAccount.Balance)
}