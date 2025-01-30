package database

type Transaction struct {
	ID        uint64  `json:"id" gorm:"primary_key"`
	VoucherNo string  `json:"voucherNo" gorm:"not null"`
	Date      string  `json:"date" gorm:"not null"`
	Amount    float64 `json:"amount" gorm:"not null"`
	Contact   string  `json:"contact" gorm:"not null"`
	Cashbook  uint64  `json:"cashbook" gorm:"not null"`
	Reference string  `json:"reference" gorm:"not null"`
	Remarks   string  `json:"remarks" gorm:"not null"`
	Category  string  `json:"category" gorm:"not null"`
}
