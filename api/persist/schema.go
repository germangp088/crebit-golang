package schema

import (
	balancedb "crebit-golang/api/persist/balancedb"
	dbhelper "crebit-golang/api/persist/dbhelper"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Init() {
	database := dbhelper.OpenDb()
	defer database.Close()
	CreateTransaction(database)
	CreateBalance(database)
}

func CreateTransaction(database *sql.DB) {
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS transactions (transaction_id INTEGER PRIMARY KEY, ttype TEXT, amount NUMERIC, effectiveDate TEXT)")
	statement.Exec()
}

func CreateBalance(database *sql.DB) {
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS balance (balance_id INTEGER PRIMARY KEY, amount NUMERIC)")
	statement.Exec()
	balancedb.CreateAmount(0)
}
