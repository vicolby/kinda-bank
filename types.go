package main

import (
	"math/rand"
	"time"
)

type TransferRequest struct {
	ToAccountID int     `json:"toAccountId"`
	Amount      float64 `json:"amount"`
}
type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Account struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	Number     int64     `json:"number"`
	Balance    float64   `json:"balance"`
	Created_at time.Time `json:"created_at"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName:  firstName,
		LastName:   lastName,
		Number:     int64(rand.Intn(1000000000)),
		Created_at: time.Now().UTC(),
	}
}
