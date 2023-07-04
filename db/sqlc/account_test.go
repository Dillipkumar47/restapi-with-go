package db

import (
	"context"
	"database/sql"
	"simplebank/db/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	args := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

}

func CreateRandomAccount(t *testing.T) Account {

	args := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account

}

func TestGetAccount(t *testing.T) {
	account := CreateRandomAccount(t)
	account1, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account1)

	require.Equal(t, account1.ID, account.ID)
	require.Equal(t, account1.Balance, account.Balance)
	require.Equal(t, account1.Owner, account.Owner)
	require.Equal(t, account1.Currency, account.Currency)
	require.Equal(t, account1.CreatedAt, account.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	accountExp := CreateRandomAccount(t)
	args := UpdateAccountParams{
		ID:      accountExp.ID,
		Balance: util.RandomBalance(),
	}
	accountAct, err := testQueries.UpdateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, accountAct)

	require.Equal(t, accountExp.ID, accountAct.ID)
	require.Equal(t, args.Balance, accountAct.Balance)
	require.Equal(t, accountExp.Owner, accountAct.Owner)
	require.Equal(t, accountExp.Currency, accountAct.Currency)
	require.Equal(t, accountExp.CreatedAt, accountAct.CreatedAt)
}

func TestDeleteAccount(t *testing.T) {
	account := CreateRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)

	require.NoError(t, err)

	account1, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.Empty(t, account1)
	require.EqualError(t, err, sql.ErrNoRows.Error())

}

func TestListAccounts(t *testing.T) {

	for i := 0; i < 10; i++ {
		CreateRandomAccount(t)
	}

	args := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, accounts, 5)
	for _, account := range accounts {
		require.NotEmpty(t, account)

	}

}
