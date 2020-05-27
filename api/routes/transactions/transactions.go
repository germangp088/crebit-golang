package transactions

import (
	balance "crebit-golang/api/models/balance"
	transaction "crebit-golang/api/models/transaction"
	transactionTypes "crebit-golang/api/types/transaction"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GetTransactions Get all transactions
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	ts := transaction.GetTransactions()
	json.NewEncoder(w).Encode(ts)
}

//GetTransaction Get a transaction by id
func GetTransactionById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	i, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal("Invalid param id")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	t := transaction.GetTransactionById(i)
	if t.TransactionId == 0 {
		log.Println("Not found")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(nil)
		return
	}
	json.NewEncoder(w).Encode(t)
}

//PostTransaction Add a new transaction
func PostTransaction(w http.ResponseWriter, r *http.Request) {

	var t transactionTypes.Transaction
	_ = json.NewDecoder(r.Body).Decode(&t)

	_, err := balance.UpdateAmount(t.Amount, t.TType)

	if err != nil {
		w.WriteHeader(409)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	id := transaction.CreateTransaction(t.Amount, t.TType)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)
}
