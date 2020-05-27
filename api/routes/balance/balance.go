package balance

import (
	balance "crebit-golang/api/models/balance"
	"encoding/json"
	"net/http"
)

//GetBalance Get balance
func GetBalance(w http.ResponseWriter, r *http.Request) {
	b := balance.GetBalance()
	json.NewEncoder(w).Encode(b)
}
