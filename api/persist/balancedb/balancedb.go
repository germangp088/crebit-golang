package balancedb

import (
	dbhelper "crebit-golang/api/persist/dbhelper"
	balance "crebit-golang/api/types/balance"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func GetAmountById(id int) balance.Balance {
	println("GetAmountById: ", id)

	database := dbhelper.OpenDb()
	defer database.Close()

	var stmt, _ = database.Prepare("SELECT balance_id, amount FROM balance WHERE balance_id = ?")
	defer stmt.Close()

	var balance_id int
	var amount float32

	err := stmt.QueryRow(strconv.Itoa(id)).Scan(&balance_id, &amount)
	if err != nil {
		log.Fatal(err)
	}

	println("balance_id: ", strconv.Itoa(balance_id))
	println("amount: ", amount)
	return balance.Balance{BalanceId: balance_id, Amount: amount}
}

func UpdateAmount(amount float32, balance_id int) {
	println("UpdateAmount: ", amount, balance_id)

	database := dbhelper.OpenDb()
	defer database.Close()

	var statement, _ = database.Prepare("UPDATE balance SET amount = ? WHERE balance_id = ?;")
	statement.Exec(amount, balance_id)
}

func CreateAmount(amount float32) {
	println("CreateAmount: ", amount)

	database := dbhelper.OpenDb()
	defer database.Close()

	var statement, _ = database.Prepare("INSERT INTO balance (amount) VALUES (?);")
	statement.Exec(amount)
}

func GetID() int {
	println("GetID")

	database := dbhelper.OpenDb()
	defer database.Close()

	rows, _ := database.Query("SELECT balance_id FROM balance")
	var balance_id int

	for rows.Next() {
		rows.Scan(&balance_id)
	}

	println("balance_id: ", balance_id)
	return balance_id
}
