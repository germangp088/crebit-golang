package transaction

type Transaction struct {
	TransactionId int     `json:"transaction_id,omitempty"`
	TType         string  `json:"type,omitempty"`
	Amount        float32 `json:"amount,omitempty"`
	EffectiveDate string  `json:"effectiveDate,omitempty"`
}

type Transactions []*Transaction
