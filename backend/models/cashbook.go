package models

type CashbookSingle struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"-"`
	TotalIncoming float64 `json:"totalIn"`
	TotalOutgoing float64 `json:"totalOut"`
	TotalBalance  float64 `json:"netBalance"`
	Currency      string  `json:"-"`
}

type CreateCashBookInput struct {
	Name           string `json:"name" binding:"required"`
	UserId         string `json:"userId" binding:"required"`
	OrganizationId string `json:"organizationId"`
}

// type Cash
