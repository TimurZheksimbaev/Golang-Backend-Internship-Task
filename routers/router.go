package routers

import (
	"go-backend-internship/handlers"
	"github.com/gorilla/mux"
)


func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/accounts", handlers.CreateAccountHandler).Methods("POST")
	r.HandleFunc("/accounts/{id}/deposit", handlers.DepositHandler).Methods("POST")
	r.HandleFunc("/accounts/{id}/withdraw", handlers.WithdrawHandler).Methods("POST")
	r.HandleFunc("/accounts/{id}/balance", handlers.BalanceHandler).Methods("GET")
	return r
}