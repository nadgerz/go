package pae

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10.0))

		got := wallet.Balance()

		// fmt.Printf("address of balance in test    is %v\n", &wallet.balance)

		want := Bitcoin(10.0)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{
			balance: Bitcoin(20.0),
		}

		wallet.Withdraw(Bitcoin(10.0))

		got := wallet.Balance()

		want := Bitcoin(10.0)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

}
