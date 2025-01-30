package database

import "gorm.io/datatypes"

type Transaction struct {
	ID        string         `json:"id" gorm:"primary_key"`
	VoucherNo string         `json:"voucherNo" gorm:"not null"`
	Date      datatypes.Date `json:"date" gorm:"not null"`
	Amount    float64        `json:"amount" gorm:"not null"`
	Contact   string         `json:"contact" gorm:"not null"`
	Cashbook  string         `json:"cashbook" gorm:"not null"`
	Reference string         `json:"reference" gorm:"not null"`
	Remarks   string         `json:"remarks" gorm:"not null"`
	Category  string         `json:"category" gorm:"not null"`
}
