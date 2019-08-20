package pae

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertError := func(t *testing.T, err error) {
		t.Helper()

		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	}

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()

		// fmt.Printf("address of balance in test    is %v\n", &wallet.balance)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10.0))

		assertBalance(t, wallet, Bitcoin(10.0))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{
			balance: Bitcoin(20.0),
		}

		wallet.Withdraw(Bitcoin(10.0))

		assertBalance(t, wallet, Bitcoin(10.0))
	})

	t.Run("Withdraw insufficent funds", func(t *testing.T) {
		startingBalance := Bitcoin(20.0)

		wallet := Wallet{startingBalance}

		err := wallet.Withdraw(Bitcoin(100.0))

		assertBalance(t, wallet, startingBalance) // nothing should change

		assertError(t, err)
	})

}
