package payment

import "errors"

type PaymentDetails struct {
    Amount   float64
    Currency string
    Method   string
}

type PaymentGateway interface {
    Validate(details PaymentDetails) error
    Process(details PaymentDetails) error
}

type BaseGateway struct{}

func (b *BaseGateway) Validate(details PaymentDetails) error {
    if details.Amount <= 0 {
        return errors.New("invalid amount")
    }
    if details.Currency == "" {
        return errors.New("currency is required")
    }
    return nil
}

type PayPalGateway struct {
    BaseGateway
}

func (p *PayPalGateway) Process(details PaymentDetails) error {
    return nil
}

type StripeGateway struct {
    BaseGateway
}

func (s *StripeGateway) Process(details PaymentDetails) error {
    return nil
}

type PaymentProcessor struct {
    gateways map[string]PaymentGateway
}

func NewPaymentProcessor() *PaymentProcessor {
    return &PaymentProcessor{
        gateways: map[string]PaymentGateway{
            "paypal": &PayPalGateway{},
            "stripe": &StripeGateway{},
        },
    }
}

func (p *PaymentProcessor) ProcessPayment(details PaymentDetails) error {
    gateway, exists := p.gateways[details.Method]
    if !exists {
        return errors.New("unsupported payment method")
    }

    if err := gateway.Validate(details); err != nil {
        return err
    }

    return gateway.Process(details)
}
