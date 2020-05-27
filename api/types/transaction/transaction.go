package transaction

type Transaction struct {
	TransactionId int     `json:"transaction_id"`
	TType         string  `json:"type"`
	Amount        float32 `json:"amount"`
	EffectiveDate string  `json:"effectiveDate"`
}

type Transactions []*Transaction
