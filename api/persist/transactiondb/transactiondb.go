package transactiondb

import (
	dbhelper "crebit-golang/api/persist/dbhelper"
	"crebit-golang/api/types/transaction"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTransaction(amount float32, ttype string, effectiveDate string) int64 {
	println("CreateAmount: ", amount)

	database := dbhelper.OpenDb()
	defer database.Close()

	statement, _ := database.Prepare("INSERT INTO transactions (amount, ttype, effectiveDate) VALUES (?, ?, ?);")
	result, _ := statement.Exec(amount, ttype, effectiveDate)
	lastID, err := result.LastInsertId()
	if err != nil {
		log.Panic(err)
	}
	return lastID
}

func GetTransactions() []*transaction.Transaction {
	println("GetTransactions")

	database := dbhelper.OpenDb()
	defer database.Close()

	rows, _ := database.Query("SELECT transaction_id, ttype, amount, effectiveDate FROM transactions;")

	transactions := []*transaction.Transaction{}
	var transaction_id int
	var ttype string
	var amount float32
	var effectiveDate string

	for rows.Next() {
		rows.Scan(&transaction_id)
		rows.Scan(&ttype)
		rows.Scan(&amount)
		rows.Scan(&effectiveDate)
		transactions = append(transactions, &transaction.Transaction{TransactionId: transaction_id, TType: ttype, Amount: amount, EffectiveDate: effectiveDate})
	}

	return transactions
}

func GetTransactionById(id int) transaction.Transaction {
	println("GetTransactionById: ", id)

	database := dbhelper.OpenDb()
	defer database.Close()

	var stmt, _ = database.Prepare("SELECT transaction_id, ttype, amount, effectiveDate FROM transactions WHERE transaction_id = ?")
	defer stmt.Close()

	var transaction_id int
	var ttype string
	var amount float32
	var effectiveDate string

	stmt.QueryRow(strconv.Itoa(id)).Scan(&transaction_id, &ttype, &amount, &effectiveDate)

	println(strconv.Itoa(transaction_id))
	println(ttype)
	println(amount)
	println(effectiveDate)

	return transaction.Transaction{TransactionId: transaction_id, TType: ttype, Amount: amount, EffectiveDate: effectiveDate}
}
