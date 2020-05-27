package transactiondb

import (
	transaction "crebit-golang/api/models/transaction"
	dbhelper "crebit-golang/api/persist/dbhelper"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTransaction(amount float32, ttype string, effectiveDate string) {
	println("CreateAmount: ", amount)

	database = dbhelper.OpenDb()
	defer database.Close()

	statement, _ = database.Prepare("INSERT INTO transaction (amount, type, effectiveDate) VALUES (?, ?, ?);")
	statement.Exec(amount, ttype, effectiveDate)
}

func GetTransactions() {
	println("GetTransactions")

	database = dbhelper.OpenDb()
	defer database.Close()

	rows, _ := database.Query("SELECT transaction_id, ttype, amount, effectiveDate FROM transaction;")

	var transactions []transaction.Transactions
	var transaction_id int
	var ttype string
	var amount float32
	var effectiveDate string

	for rows.Next() {
		rows.Scan(&transaction_id)
		rows.Scan(&ttype)
		rows.Scan(&amount)
		rows.Scan(&effectiveDate
		transactions = append(transactions, transaction.Transaction{TransactionId: transaction_id, TType: ttype, Amount: amount, EffectiveDate: effectiveDate}))
	}

	println("transactions:", transactions)
	return transactions
}

func GetTransactionById(transaction_id int) {
	println("GetTransactionById: ", transaction_id)

	database = dbhelper.OpenDb()
	defer database.Close()

	rows, _ := database.Query("SELECT transaction_id, type, amount, effectiveDate FROM transaction WHERE transaction_id = ?;", transaction_id)
	var transaction_id int
	var ttype string
	var amount float32
	var effectiveDate string

	for rows.Next() {
		rows.Scan(&transaction_id)
		rows.Scan(&ttype)
		rows.Scan(&amount)
		rows.Scan(&effectiveDate
	}

	println(strconv.Itoa(transaction_id))
	println(strconv.Itoa(ttype))
	println(strconv.Itoa(amount))
	println(strconv.Itoa(effectiveDate))
	return transaction.Transaction{TransactionId: transaction_id, TType: ttype, Amount: amount, EffectiveDate: effectiveDate}
}
