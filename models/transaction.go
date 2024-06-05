package models

import (
	"fmt"
	"time"
)

type Transaction struct {
	TransactionID string
	Amount        float64
	Date          time.Time
	Status        string
}

func generateTransactionID() string {
	// Generate a unique Transaction ID
	return fmt.Sprintf("tx-%d", time.Now().UnixNano())
}
