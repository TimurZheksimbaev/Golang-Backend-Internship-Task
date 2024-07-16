package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
	"go-backend-internship/accounts"
	"math/rand"
)

var accountMap = make(map[string]*accounts.Account)
var accountMapMu sync.Mutex
const ID_LENGTH int = 5


func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	account := accounts.NewAccount(randomID())
	accountMapMu.Lock()
	accountMap[account.ID] = account												
	accountMapMu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": account.ID})
}


func DepositHandler(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id := vars.Get("id")
	amount := vars.Get("amount")
	if id == "" || amount == "" {
		http.Error(w, "missing id or amount", http.StatusBadRequest)
		return
	}
	accountMapMu.Lock()
	account, exists := accountMap[id]
	accountMapMu.Unlock()
	if !exists {
		http.Error(w, "account not found", http.StatusNotFound)
		return
	}
	amountFloat := parseAmount(amount)
	if amountFloat <= 0 {
		http.Error(w, "invalid amount", http.StatusBadRequest)
		return
	}
	account.Deposit(amountFloat)
	w.WriteHeader(http.StatusOK)
}


func WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id := vars.Get("id")
	amount := vars.Get("amount")
	if id == "" || amount == "" {
		http.Error(w, "missing id or amount", http.StatusBadRequest)
		return
	}
	accountMapMu.Lock()
	account, exists := accountMap[id]
	accountMapMu.Unlock()
	if !exists {
		http.Error(w, "account not found", http.StatusNotFound)
		return
	}
	amountFloat := parseAmount(amount)
	if amountFloat <= 0 {
		http.Error(w, "invalid amount", http.StatusBadRequest)
		return
	}
	account.Withdraw(amountFloat)
	w.WriteHeader(http.StatusOK)
}

// balanceHandler handles balance requests
func BalanceHandler(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id := vars.Get("id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}
	accountMapMu.Lock()
	account, exists := accountMap[id]
	accountMapMu.Unlock()
	if !exists {
		http.Error(w, "account not found", http.StatusNotFound)
		return
	}
	balance := account.GetBalance()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"balance": balance})
}

func parseAmount(s string) float64 {
	var f float64
	fmt.Sscanf(s, "%f", &f)
	return f
}


func randomID() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, ID_LENGTH)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}