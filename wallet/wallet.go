package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// Deposit receive wallet balance, accept deposit amount
func (w *Wallet) Deposit(amount Bitcoin) {
	// combine balance to amount received
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
