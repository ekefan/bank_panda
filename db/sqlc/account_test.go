package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ekefan/bank_panda/utils"
	"github.com/stretchr/testify/require"
)

//create the params for creating an account

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner: user.Username,
		Balance: utils.RandomBalance(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

//function to test the creation of an account
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t);
}


func TestDeleteAccount(t *testing.T) {
	testDelAcc := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), testDelAcc.ID)
	require.NoError(t, err)
	
	checkAcc, err := testQueries.GetAccount(context.Background(), testDelAcc.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, checkAcc)
}


func TestGetAccount(t *testing.T) {
	acc := createRandomAccount(t)
	accGotten, err := testQueries.GetAccount(context.Background(), acc.ID)
	
	require.NoError(t, err)
	require.NotEmpty(t, accGotten)

	require.Equal(t, acc.ID, accGotten.ID)
	require.Equal(t, acc.Owner, accGotten.Owner)
	require.Equal(t, acc.Balance, accGotten.Balance)
	require.Equal(t, acc.Currency, accGotten.Currency)
	require.Equal(t, acc.CreatedAt, accGotten.CreatedAt)
	require.WithinDuration(t, acc.CreatedAt, accGotten.CreatedAt, time.Second)
}


func TestUpdateAccount(t *testing.T) {
	acc := createRandomAccount(t)
	newBalance := utils.RandomBalance()

	updateParams := UpdateAccountParams{
		ID: acc.ID,
		Balance: newBalance,
	}
	updatedAcc, err := testQueries.UpdateAccount(context.Background(), updateParams)
	
	require.NoError(t, err)
	require.NotEmpty(t, updatedAcc)
	require.NotEqual(t, acc.Balance, updatedAcc.Balance)
	require.Equal(t, newBalance, updatedAcc.Balance)

}


func TestListAccounts(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	
	}
	listParams := ListAccountsParams{
		Owner: lastAccount.Owner,		
		Limit: 5,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), listParams)

	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, acc := range accounts {
		require.NotEmpty(t, acc)
		require.Equal(t, lastAccount.Owner, acc.Owner)
	}

}