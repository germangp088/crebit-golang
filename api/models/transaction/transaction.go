package transaction

import (
	transactiondb "crebit-golang/api/persist/transactiondb"
	transaction "crebit-golang/api/types/transaction"
	"time"
)

func GetTransactions() []transaction.Transaction {
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

func CreateTransaction(amount float32, ttype string) int64 {
	println("CreateTransaction:")
	println("ttype:", ttype)
	println("amount:", amount)

	effectiveDate := time.Now().UTC().Format(time.RFC3339)

	id := transactiondb.CreateTransaction(amount, ttype, effectiveDate)
	println("id:", id)
	return id
}
