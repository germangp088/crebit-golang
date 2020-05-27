package transaction

import (
	transactiondb "crebit-golang/api/persist/transactiondb"
	transaction "crebit-golang/api/types/transaction"
)

func GetTransactions() []*transaction.Transaction {
	println("GetTransactions")
	transactions := transactiondb.GetTransactions()
	return transactions
}
