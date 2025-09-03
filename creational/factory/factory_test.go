package creational

import (
	"strings"
	"testing"
)

func TestPaymentMethodExists(t *testing.T){

	cardPaymentMethod, err := GetPaymentMethod(CardType)
	if err != nil {
		t.Fatal("a payment method of type 'Card' must exist")
	}

	bankPaymentMethod, err := GetPaymentMethod(BankTransferType)
	if err != nil {
		t.Fatal("a payment method of type 'BankTransfer' must exist")
	}

	cardResult := cardPaymentMethod.Pay(100)
	bankResult := bankPaymentMethod.Pay(100)

	if strings.Contains(cardResult, "paid with credit card") == false {
		t.Fatal("card payment method must contain the string 'paid with credit card'")
	}
	if strings.Contains(bankResult, "paid with bank transfer") == false {
		t.Fatal("bank payment method must contain the string 'paid with bank transfer'")
	}

}

func TestPaymentMethodInvalid(t* testing.T) {
	_, err := GetPaymentMethod("invalid")
	if err == nil {
		t.Fatal("a payment method of type 'invalid' must not exist")
	}
}