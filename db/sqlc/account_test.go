package db

import (
	"context"
	"testing"
	"time"

	"github.com/riad/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
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

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	first_account := createRandomAccount(t)
	second_account, err := testQueries.GetAccount(context.Background(), first_account.ID)

	require.NoError(t, err)
	require.NotEmpty(t, second_account)

	require.Equal(t, first_account.ID, second_account.ID)
	require.Equal(t, first_account.Owner, second_account.Owner)
	require.Equal(t, first_account.Balance, second_account.Balance)
	require.Equal(t, first_account.Currency, second_account.Currency)
	require.WithinDuration(t, first_account.CreatedAt, second_account.CreatedAt, 2*time.Second)
}

func TestUpdateAccount(t *testing.T) {
	first_account := createRandomAccount(t)

	account_arg := UpdateAccountParams{
		ID:      first_account.ID,
		Balance: util.RandomMoney(),
	}
	updated_account, err := testQueries.UpdateAccount(context.Background(), account_arg)

	require.NoError(t, err)
	require.NotEmpty(t, updated_account)

	require.Equal(t, first_account.ID, updated_account.ID)
	require.Equal(t, first_account.Owner, updated_account.Owner)
	require.NotEqual(t, first_account.Balance, updated_account.Balance)
	require.Equal(t, first_account.Currency, updated_account.Currency)
	require.WithinDuration(t, first_account.CreatedAt, updated_account.CreatedAt, 2*time.Second)
}

func TestDeleteAccount(t *testing.T) {
	first_account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), first_account.ID)

	require.NoError(t, err)

	second_account, err := testQueries.GetAccount(context.Background(), first_account.ID)
	require.Error(t, err)
	require.Empty(t, second_account)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	list_accounts_arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), list_accounts_arg)

	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}

}
