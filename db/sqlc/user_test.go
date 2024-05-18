package db

import (
	"testing"
	"context"
	"github.com/ekefan/bank_panda/utils"
	"time"
	"github.com/stretchr/testify/require"
)


//create the params for creating an account

func createRandomUser(t *testing.T) User {
	hashPassword , err := utils.HashPassword(utils.RandomString(6))
	require.NoError(t, err)
	
	arg := CreateUserParams{
		Username: utils.RandomOwner(),
		HashedPassword: hashPassword,
		FullName:  utils.RandomOwner(),
		Email: utils.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

//function to test the creation of an account
func TestCreateUser(t *testing.T) {
	createRandomUser(t);
}


// func TestDeleteUser(t *testing.T) {
// 	testDelUsr := createRandomUser(t)
// 	err := testQueries.DeleteAccount(context.Background(), testDelAcc.ID)
// 	require.NoError(t, err)
	
// 	checkAcc, err := testQueries.GetAccount(context.Background(), testDelAcc.ID)

// 	require.Error(t, err)
// 	require.EqualError(t, err, sql.ErrNoRows.Error())
// 	require.Empty(t, checkAcc)
// }


func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}