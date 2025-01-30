package database

type CashFlow struct {
	ID        uint64  `json:"id"`
	Amount    float64 `json:"amount"`
	Direction string  `json:"direction"`
	Note      string  `json:"note"`
	Cashbook  uint64  `json:"cashbook"`
}
