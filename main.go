package main

import (
    "fmt"
    "payment-processing/payment"
)

func main() {
    processor := payment.NewPaymentProcessor()

    // Example usage
    details := payment.PaymentDetails{
        Amount:   100.0,
        Currency: "USD",
        Method:   "paypal",
    }

    err := processor.ProcessPayment(details)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Payment processed successfully!")
    }
}
