package pae

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10.0))

		got := wallet.Balance()

		// fmt.Printf("address of balance in test    is %v\n", &wallet.balance)

		want := Bitcoin(10.0)

		if got != want {
			t.Errorf("%#v got %s want %s", want, got, want)
		}

	})
}
