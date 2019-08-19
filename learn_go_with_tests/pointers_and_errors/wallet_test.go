package pae

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet()

		wallet.Deposit(10)

		got := wallet.Balance()
		want := 10

		if got != want {
			t.Errorf("%#v got %.2f want %.2f", want, got, want)
		}

	})
}
