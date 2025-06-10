package main

import "errors"

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

func updateBalance(cust *customer, trans transaction) error {
	switch trans.transactionType {
	case transactionDeposit:
		(*cust).balance += trans.amount
	case transactionWithdrawal:
		if trans.amount > (*cust).balance {
			return errors.New("insufficient funds")
		}
		(*cust).balance -= trans.amount
	default:
		return errors.New("unknown transaction type")
	}
	return nil
}

// Don't touch above this line

// ?
// alice := customer{id: 1, balance: 100.0}
// deposit := transaction{customerID: 1, amount: 50, transactionType: transactionDeposit}

// updateBalance(&alice, deposit)
