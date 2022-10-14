package account

import "sync"

// Define the Account type here.
type Account struct {
	amount int64
	closed bool
	mu     sync.Mutex
}

func Open(amount int64) *Account {
	var a *Account
	if amount >= 0 {
		a = &Account{amount: amount, closed: false}
		a.mu.Lock()
		defer a.mu.Unlock()
	}
	return a
}

func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.closed {
		return 0, false
	}
	return a.amount, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	switch {
	case a.closed:
		return 0, false
	case a.amount+amount < 0:
		return a.amount, false
	default:
		a.amount += amount
		return a.amount, true
	}
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.closed {
		return 0, false
	}
	a.closed = true
	b := a.amount
	a.amount = 0
	return b, true
}
