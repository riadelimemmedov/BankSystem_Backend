package db

import (
	"context"
	"fmt"
	"testing"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	first_account := createRandomAccount(t)
	second_account := createRandomAccount(t)

	n := 5
	amount := int64(10)

	// errs := make(chan error)
	// results := make(chan TransferTxResult)

	//!Run give times this loop,each time create seperate go routine for handling concurrent transaction,this prevent single point of failure and head of line blocing issue
	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: first_account.ID,
				ToAccountID:   second_account.ID,
				Amount:        amount,
			})
			fmt.Println("Result", result)
			fmt.Println("Err", err)

			// errs <- err
			// results <- result
		}()
	}

	// //!Loop all transfers come from go channels
	// for i := 0; i < n; i++ {
	// 	err := <-errs
	// 	require.NoError(t, err)

	// 	result := <-results
	// 	require.NotEmpty(t, result)

	// 	transfer := result.Transfer

	// 	//!Check Transfer
	// 	require.NotEmpty(t, transfer)
	// 	require.Equal(t, first_account.ID, transfer.FromAccountID)
	// 	require.Equal(t, second_account.ID, transfer.ToAccountID)
	// 	require.Equal(t, amount, transfer.Amount)
	// 	require.NotZero(t, transfer.ID)
	// 	require.NotZero(t, transfer.CreatedAt)

	// 	//!Get created transfer detail,we check transfer successfully saved to db or not
	// 	_, err = store.GetTransfer(context.Background(), transfer.ID)
	// 	require.NoError(t, err)

	// 	//!Check FromEntry
	// 	from_entry := result.FromEntry
	// 	require.NotEmpty(t, from_entry)
	// 	require.Equal(t, first_account.ID, from_entry.AccountID)
	// 	require.Equal(t, -amount, from_entry.Amount)
	// 	require.NotZero(t, from_entry.ID)
	// 	require.NotZero(t, from_entry.CreatedAt)

	// 	//!Check entry successfully created or not
	// 	_, err = store.GetEntry(context.Background(), from_entry.ID)
	// 	require.NoError(t, err)

	// 	//!Check ToEntry
	// 	to_entry := result.ToEntry
	// 	require.NotEmpty(t, to_entry)
	// 	require.Equal(t, second_account.ID, to_entry.AccountID)
	// 	require.Equal(t, amount, to_entry.Amount)
	// 	require.NotZero(t, to_entry.ID)
	// 	require.NotZero(t, to_entry.CreatedAt)

	// 	//!Check entry successfully created or not
	// 	_, err = store.GetEntry(context.Background(), to_entry.ID)
	// 	require.NoError(t, err)
	// }
}
