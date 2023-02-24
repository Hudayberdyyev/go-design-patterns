package factory

import (
	"errors"
	"fmt"
)

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

type CreditCardPM struct {
}

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(CreditCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("payment method %d not recognized", m))
	}
}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

func (d *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card\n", amount)
}

func (c *CreditCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using debit card (new)\n", amount)
}
