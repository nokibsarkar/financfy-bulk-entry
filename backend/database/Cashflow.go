package database

type CashFlow struct {
	ID            uint64   `json:"id"`
	Date          string   `json:"date"`
	Amount        float64  `json:"amount"`
	TotalIncoming float64  `json:"totalIncoming"`
	TotalOutgoing float64  `json:"totalOutgoing"`
	TotalBalance  float64  `json:"totalBalance"`
	CashbookID    uint64   `json:"cashbookId"`
	Cashbook      Cashbook `json:"-" gorm:"foreignKey:CashbookID"`
}
