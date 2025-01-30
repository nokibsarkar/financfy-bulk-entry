package database

type Cashbook struct {
	ID            uint64  `json:"id" gorm:"primary_key"`
	Name          string  `json:"name" gorm:"not null"`
	Description   string  `json:"-" gorm:"null"`
	TotalIncoming float64 `json:"totalIn" gorm:"default:0"`
	TotalOutgoing float64 `json:"totalOut" gorm:"default:0"`
	TotalBalance  float64 `json:"netBalance" gorm:"default:0"`
	Currency      string  `json:"-" gorm:"default:'BDT'"`
}
