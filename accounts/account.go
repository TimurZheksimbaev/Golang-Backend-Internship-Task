package accounts

import (
	"go-backend-internship/utils"
	"sync"
	"time"
)

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

type Account struct {
	ID      string
	Balance float64
	mux      sync.Mutex
	depositChan chan float64
	withdrawChan chan float64
}

func NewAccount(id string) *Account {
	a := &Account{
		ID:          id,
		Balance:     0,
		depositChan: make(chan float64),
		withdrawChan: make(chan float64),
	}
	go a.processTransactions()
	return a
}

func (a *Account) processTransactions() {
	start := time.Now()
	for {
		select {
		case amount := <-a.depositChan:
			a.mux.Lock()
			a.Balance += amount
			utils.LogInfo("Deposit", a.ID, time.Since(start))
			a.mux.Unlock()
		case amount := <-a.withdrawChan:
			a.mux.Lock()
			if a.Balance < amount {
				utils.LogInsufficientFunds(a.ID, time.Since(start), a.Balance)
			} else {
				a.Balance -= amount
				utils.LogInfo("Deposit", a.ID, time.Since(start))
			}
			a.mux.Unlock()
		}
	}
}


func (a *Account) Deposit(amount float64) error {
	a.depositChan <- amount
	return nil
}


func (a *Account) Withdraw(amount float64) error {
	a.withdrawChan <- amount
	return nil
}


func (a *Account) GetBalance() float64 {
	start := time.Now()
	a.mux.Lock()
	defer a.mux.Unlock()
	utils.LogInfo("Get Balance", a.ID, time.Since(start))
	return a.Balance
}

