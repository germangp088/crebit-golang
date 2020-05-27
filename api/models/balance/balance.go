package balance

import (
	balancedb "crebit-golang/api/persist/balancedb"
	balance "crebit-golang/api/types/balance"
	"errors"
)

func GetBalance() balance.Balance {
	println("GetBalance")
	balance_id := balancedb.GetID()
	balance := balancedb.GetAmountById(balance_id)
	println("balance:", balance.BalanceId)
	return balance
}

func UpdateAmount(amount float32, ttype string) (balance.Balance, error) {
	println("UpdateAmount")
	println("amount:", amount)
	println("ttype:", ttype)
	balance := GetBalance()

	if balance.Amount == 0 && ttype == "debit" {
		return balance, errors.New("The amount cannot be negative")
	}

	var newAmount float32
	if ttype == "credit" {
		newAmount = balance.Amount + amount
	} else {
		newAmount = balance.Amount - amount
	}

	if newAmount < 0 {
		return balance, errors.New("The amount cannot be negative")
	}

	balancedb.UpdateAmount(balance.BalanceId, newAmount)

	balance.Amount = newAmount

	println("balance:", balance.Amount)
	return balance, nil
}
