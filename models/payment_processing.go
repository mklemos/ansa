package models

import (
    "fmt"
    "time"
)

type PaymentProcessing struct {
    TransactionHistory []Transaction
}

func (p *PaymentProcessing) CreateChargeRequest(customer *CustomerAccount, amount float64) string {
    if customer.PreloadedBalance < amount {
        return "Insufficient funds"
    }

    customer.DeductFunds(amount)
    transaction := Transaction{
        TransactionID: generateTransactionID(),
        Amount:        amount,
        Date:          time.Now(),
        Status:        "Charged",
    }
    p.TransactionHistory = append(p.TransactionHistory, transaction)
    return fmt.Sprintf("Charge request created for customer %s for amount $%.2f", customer.CustomerID, amount)
}

func (p *PaymentProcessing) HandlePaymentResponse(customer *CustomerAccount, success bool, amount float64) string {
    if success {
        customer.AddFunds(amount)
        return fmt.Sprintf("Payment of $%.2f was successful for customer %s", amount, customer.CustomerID)
    } else {
        customer.RefundFunds(amount)
        return fmt.Sprintf("Payment of $%.2f failed for customer %s", amount, customer.CustomerID)
    }
}

func (p *PaymentProcessing) ProcessRefund(customer *CustomerAccount, amount float64) string {
    customer.RefundFunds(amount)
    transaction := Transaction{
        TransactionID: generateTransactionID(),
        Amount:        amount,
        Date:          time.Now(),
        Status:        "Refunded",
    }
    p.TransactionHistory = append(p.TransactionHistory, transaction)
    return fmt.Sprintf("Refund of $%.2f processed for customer %s", amount, customer.CustomerID)
}
