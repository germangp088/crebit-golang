package balance

import (
	balancedb "crebit-golang/api/persist/balancedb"
	balance "crebit-golang/api/types/balance"
)

func GetBalance() balance.Balance {
	println("GetBalance")
	balance_id := balancedb.GetID()
	balance := balancedb.GetAmountById(balance_id)
	println("balance:", balance.BalanceId)
	return balance
}
