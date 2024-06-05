package models

type PaymentRequest struct {
    CustomerID string  `json:"customer_id"`
    Amount     float64 `json:"amount"`
}

type PaymentResponse struct {
    Message string `json:"message"`
}
