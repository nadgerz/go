package pae

import (
	"testing"
)

func TestWallet(t *testing.T) {
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

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20.0)
		wallet := Wallet{startingBalance}

		err := wallet.Withdraw(Bitcoin(100.0))

		assertBalance(t, wallet, startingBalance) // balance should not change
		assertError(t, err, ErrInsufficientFunds)
	})

}

func assertError(t *testing.T, got error, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	// fmt.Printf("address of balance in test    is %v\n", &wallet.balance)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
