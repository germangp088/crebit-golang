package transactiondb

import (
	dbhelper "crebit-golang/api/persist/dbhelper"
	transactionTypes "crebit-golang/api/types/transaction"
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

func GetTransactions() []transactionTypes.Transaction {
	println("GetTransactions")

	database := dbhelper.OpenDb()
	defer database.Close()

	rows, _ := database.Query("SELECT transaction_id, ttype, amount, effectiveDate FROM transactions")

	defer rows.Close()

	transactions := []transactionTypes.Transaction{}

	for rows.Next() {
		var transaction_id int
		var ttype string
		var amount float32
		var effectiveDate string
		err := rows.Scan(&transaction_id, &ttype, &amount, &effectiveDate)
		if err != nil {
			log.Fatal(err)
		}
		transaction := transactionTypes.Transaction{TransactionId: transaction_id, TType: ttype, Amount: amount, EffectiveDate: effectiveDate}
		transactions = append(transactions, transaction)
	}

	return transactions
}

func GetTransactionById(id int) transactionTypes.Transaction {
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

	println("transaction_id: ", strconv.Itoa(transaction_id))
	println("ttype: ", ttype)
	println("amount: ", amount)
	println("effectiveDate: ", effectiveDate)

	return transactionTypes.Transaction{TransactionId: transaction_id, TType: ttype, Amount: amount, EffectiveDate: effectiveDate}
}
