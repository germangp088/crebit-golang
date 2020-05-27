//go:generate swagger generate spec
package balance

import (
	"encoding/json"
	"net/http"
)

//GetBalance Get balance
func GetBalance(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(0)
}
