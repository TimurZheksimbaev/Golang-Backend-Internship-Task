package utils

import (
	"log"
	"time"
)


func LogInfo(operation string, id string, execTime time.Duration) {
	operationTime := time.Now().Format("2006-01-02 15:04:05")
	log.Printf("%s Operation: %s Id: %s  Time took: %d", operationTime, operation, id, execTime)
}

func LogInsufficientFunds(id string, execTime time.Duration, balance float64) {
	operationTime := time.Now().Format("2006-01-02 15:04:05")
	log.Printf("%s Account %s has not enough funds for withdrawal. Current balance: %f Time took: %s", operationTime, id, balance, execTime)
}

