package balance

import (
	balancedb "crebit-golang/api/persist/balancedb"
	balance "crebit-golang/api/types/balance"
)

func GetBalance() balance.Balance {
	balance_id := balancedb.GetID()
	var balance = balancedb.GetAmountById(balance_id)
	return balance
}
