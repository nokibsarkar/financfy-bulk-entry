package cashbook

type Cashbook struct{}

type ICashbookService interface {
	CreateCashbook(cashbook *Cashbook) error
	GetCashbook(cashbookID string) (*Cashbook, error)
}

func (c *Cashbook) CreateCashbook(cashbook *Cashbook) error {
	return nil
}

func (c *Cashbook) GetCashbook(cashbookID string) (*Cashbook, error) {
	return nil, nil
}
