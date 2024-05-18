package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)


func TestPassword(t *testing.T) {
	password := RandomString(8)
	fmt.Println(password)
	hash1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hash1)
	

	err = CheckPassword(password, hash1)
	require.NoError(t, err)


	wrongPassword := RandomString(6)
	err = CheckPassword(wrongPassword, hash1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hash2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hash1)
	require.NotEqual(t,hash1, hash2)
}
