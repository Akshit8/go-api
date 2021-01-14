/*
 * @File: db.sqlc.account_test.go
 * @Description: unit tests for account operations
 * @LastModifiedTime: 2021-01-14 17:06:01
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatingAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner: "Akshit",
		Balance: 200,
		Currency: "INR",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}
