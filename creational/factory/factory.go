package factory

import "errors"

const (
	Cash      = 1
	DebitCard = 2
)

type PaymentMethod interface {
	Pay(amount float32) string
}

type CashPM struct {
}

type DebitCardPM struct {
}

func GetPaymentMethod(m int) (PaymentMethod, error) {
	return nil, errors.New("not implemented yet")
}

func (c *CashPM) Pay(amount float32) string {
	return ""
}

func (c *DebitCardPM) Pay(amount float32) string {
	return ""
}
