package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/akmal4410/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)

	arg := CreateAccountParams{
		Owner:    user.UserName,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	accout, err := testQuires.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accout)

	require.Equal(t, arg.Owner, accout.Owner)
	require.Equal(t, arg.Balance, accout.Balance)
	require.Equal(t, arg.Currency, accout.Currency)

	require.NotZero(t, accout.ID)
	require.NotZero(t, accout.CreatedAt)

	return accout
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	creAccount := createRandomAccount(t)
	getAccount, err := testQuires.GetAccount(context.Background(), creAccount.ID)

	require.NoError(t, err)
	require.NotEmpty(t, getAccount)

	require.Equal(t, creAccount.ID, getAccount.ID)
	require.Equal(t, creAccount.Owner, getAccount.Owner)
	require.Equal(t, creAccount.Balance, getAccount.Balance)
	require.Equal(t, creAccount.Currency, getAccount.Currency)
	require.WithinDuration(t, creAccount.CreatedAt, getAccount.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	creAccount := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      creAccount.ID,
		Balance: util.RandomMoney(),
	}

	updAccount, err := testQuires.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updAccount)

	require.Equal(t, creAccount.ID, updAccount.ID)
	require.Equal(t, creAccount.Owner, updAccount.Owner)
	require.Equal(t, arg.Balance, updAccount.Balance)
	require.Equal(t, creAccount.Currency, updAccount.Currency)
	require.WithinDuration(t, creAccount.CreatedAt, updAccount.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	creAccount := createRandomAccount(t)
	err := testQuires.DeleteAccount(context.Background(), creAccount.ID)

	require.NoError(t, err)
	getAccount, err := testQuires.GetAccount(context.Background(), creAccount.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, getAccount)

}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	args := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQuires.ListAccounts(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
