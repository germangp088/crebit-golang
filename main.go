package main

import (
	"log"
	"net/http"

	schema "crebit-golang/api/persist"
	balance "crebit-golang/api/routes/balance"
	transactions "crebit-golang/api/routes/transactions"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	Initialize()
}

//Initialize init of services.
func Initialize() {
	schema.Init()

	router := mux.NewRouter()

	router.HandleFunc("/transactions", transactions.GetTransactions).Methods("GET")
	router.HandleFunc("/transactions/{id}", transactions.GetTransactionById).Methods("GET")
	router.HandleFunc("/transactions", transactions.PostTransaction).Methods("POST")
	router.HandleFunc("/balance", balance.GetBalance).Methods("GET")

	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
