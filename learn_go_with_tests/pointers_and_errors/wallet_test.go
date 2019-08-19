package pae

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10.0)

		got := wallet.Balance()

		fmt.Printf("address of balance in test is    %v\n", &wallet.balance)

		want := 10.0

		if got != want {
			t.Errorf("%#v got %.2f want %.2f", want, got, want)
		}

	})
}
