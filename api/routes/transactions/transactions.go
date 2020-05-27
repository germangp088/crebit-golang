package transactions

import (
	"crebit-golang/api/models/transaction"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// HTTP status code 200 and repository model in data
type Resp struct {
	Body struct {
		// HTTP status code 200/201
		Code int `json:"code"`
		Data int `json:"id"`
	}
}

// HTTP status code 500/400
type RespError struct {
	Body struct {
		// HTTP status code 500/400
		Code int    `json:"code"`
		Data string `json:"error"`
	}
}

//GetTransactions Get all transactions
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode('')
}

//GetTransaction Get a transaction by id
func GetTransaction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	i, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal("Invalid param id")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

//PostTransaction Add a new transaction
func PostTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction transaction.Transaction
	_ = json.NewDecoder(r.Body).Decode(&transaction)
	json.NewEncoder(w).Encode(transaction)
}
