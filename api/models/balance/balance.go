package balance

type Balance struct {
	BalanceId int     `json:"balance_id,omitempty"`
	Amount    float32 `json:"amount,omitempty"`
}
