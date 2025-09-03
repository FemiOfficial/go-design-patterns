package creational

import "fmt"

type PaymentMethod interface {
	Pay(amount float64) string
}

const (
	CardType = "card"
	BankTransferType = "bank_transfer"
)

type Card struct {}
type BankTransfer struct {}

func (c *Card) Pay(amount float64) string {
	return fmt.Sprintf("%0.2f paid with credit card", amount)
}

func (b *BankTransfer) Pay(amount float64) string {
	return fmt.Sprintf("%0.2f paid with bank transfer", amount)
}

func GetPaymentMethod(m string) (PaymentMethod, error) {
	switch m {
	case CardType:
		return &Card{}, nil
	case BankTransferType:
		return new(BankTransfer), nil
	default:
		return nil, fmt.Errorf("invalid payment method")
	}
}