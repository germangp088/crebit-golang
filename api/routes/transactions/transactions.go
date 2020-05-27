package transactions

import (
	transaction "crebit-golang/api/models/transaction"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GetTransactions Get all transactions
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	transactions := transaction.GetTransactions()
	json.NewEncoder(w).Encode(transactions)
}

//GetTransaction Get a transaction by id
func GetTransaction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	i, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal("Invalid param id")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(i)
}

//PostTransaction Add a new transaction
func PostTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction = transaction.GetTransactions()
	_ = json.NewDecoder(r.Body).Decode(&transaction)
	json.NewEncoder(w).Encode(transaction)
}
