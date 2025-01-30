package models

type CashbookSingle struct {
	ID            uint64  `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"-"`
	TotalIncoming float64 `json:"totalIn"`
	TotalOutgoing float64 `json:"totalOut"`
	TotalBalance  float64 `json:"netBalance"`
	Currency      string  `json:"-"`
}

type CreateCashBookInput struct {
	Name           string `json:"name" binding:"required"`
	UserId         uint64 `json:"userId" binding:"required"`
	OrganizationId uint64 `json:"organizationId"`
}

// type Cash
