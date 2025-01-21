package payment

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestPaymentProcessor(t *testing.T) {
    processor := NewPaymentProcessor()

    tests := []struct {
        name        string
        details     PaymentDetails
        expectError bool
    }{
        {"Valid PayPal Payment", PaymentDetails{Amount: 100, Currency: "USD", Method: "paypal"}, false},
        {"Valid Stripe Payment", PaymentDetails{Amount: 200, Currency: "EUR", Method: "stripe"}, false},
        {"Invalid Payment Method", PaymentDetails{Amount: 100, Currency: "USD", Method: "cash"}, true},
        {"Zero Amount", PaymentDetails{Amount: 0, Currency: "USD", Method: "paypal"}, true},
        {"Missing Currency", PaymentDetails{Amount: 100, Currency: "", Method: "paypal"}, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := processor.ProcessPayment(tt.details)
            if tt.expectError {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
            }
        })
    }
}
