package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/MirzaKarabulut/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner: user.Username,
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := TestQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	createAcc := createRandomAccount(t)
	account1, err := TestQueries.GetAccount(context.Background(), createAcc.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account1)

	require.Equal(t, createAcc, account1)
	require.Equal(t, createAcc.ID, account1.ID)
	require.Equal(t, createAcc.Owner, account1.Owner)
	require.Equal(t, createAcc.Balance, account1.Balance)
	require.Equal(t, createAcc.Currency, account1.Currency)
	require.WithinDuration(t, createAcc.CreatedAt, account1.CreatedAt, time.Second)
}

func TestListAccount(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
	 lastAccount = createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Owner: lastAccount.Owner,
		Limit: 5,
		Offset: 0,
	}
	accounts, err := TestQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.NotZero(t, account.ID)
		require.NotZero(t, account.CreatedAt)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}

func TestUpdateAccount(t *testing.T) {
	createAcc := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID: createAcc.ID,
		Balance: util.RandomMoney(),
	}
	account1, err := TestQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account1)

	require.Equal(t, createAcc.ID, account1.ID)
	require.Equal(t, createAcc.Owner, account1.Owner)
	require.Equal(t, arg.Balance, account1.Balance)
	require.Equal(t, createAcc.Currency, account1.Currency)
	require.WithinDuration(t, createAcc.CreatedAt, account1.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	createAcc := createRandomAccount(t)
	err := TestQueries.DeleteAccount(context.Background(), createAcc.ID)
	require.NoError(t, err)

	account1, err := TestQueries.GetAccount(context.Background(), createAcc.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account1)
}