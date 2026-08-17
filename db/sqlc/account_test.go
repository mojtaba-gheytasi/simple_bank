package db

import (
	"context"
	"database/sql"
	"testing"
	"fmt"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner: "alex",
		Balance: 100,
		Currency: "EUR",
	}

	account := createTestAccount(t, arg)

	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestGetAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner: "max",
		Balance: 200,
		Currency: "EUR",
	}
	account1 := createTestAccount(t, arg)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.CreatedAt, account2.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner: "max",
		Balance: 200,
		Currency: "EUR",
	}
	account1 := createTestAccount(t, arg)
	err1 := testQueries.UpdateAccount(context.Background(), UpdateAccountParams{
		ID: account1.ID,
		Balance: 990,
	})
	require.NoError(t, err1)
	account2, err2 := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, int64(990), account2.Balance)
}

func TestDeleteAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner: "max",
		Balance: 200,
		Currency: "EUR",
	}
	account1 := createTestAccount(t, arg)

	err1 := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err1)

	account2, err2 := testQueries.GetAccount(context.Background(), account1.ID)
	require.EqualError(t, err2, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccount(t *testing.T) {
	for i := range 10 {
		arg := CreateAccountParams{
			Owner: fmt.Sprintf("max%d", i),
			Balance: 100 * int64(i),
			Currency: "EUR",
		}
		createTestAccount(t, arg)
	}

	accounts, err := testQueries.ListAccount(context.Background(), ListAccountParams{
		Limit: 5,
		Offset: 5,
	})
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func createTestAccount(t *testing.T, arg CreateAccountParams) Account {
	t.Helper()
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)

	return account
}