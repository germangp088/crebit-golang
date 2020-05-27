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

func GetTransactionById(transaction_id int) transaction.Transaction {
	println("GetTransactionById:", transaction_id)
	transaction := transactiondb.GetTransactionById(transaction_id)
	println("transaction:", transaction.TransactionId)
	return transaction
}
