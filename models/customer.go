package models

import (
    "context"
	"log"
    "project/config"
	"fmt"
)

type BalanceLedger struct {
    FundsAdded    float64
    FundsSpent    float64
    FundsRefunded float64
}

type CustomerAccount struct {
    CustomerID              string
    Name                    string
    Email                   string
    EncryptedPaymentDetails string
    BalanceLedger           BalanceLedger
    LoyaltyPoints           int
    PreloadedBalance        float64
    NFTCollection           []NFT
}

func (c *CustomerAccount) AddFunds(amount float64) {
    c.PreloadedBalance += amount
    c.BalanceLedger.FundsAdded += amount
}

func (c *CustomerAccount) DeductFunds(amount float64) {
    c.PreloadedBalance -= amount
    c.BalanceLedger.FundsSpent += amount
}

func (c *CustomerAccount) RefundFunds(amount float64) {
    c.PreloadedBalance += amount
    c.BalanceLedger.FundsRefunded += amount
}

func (c *CustomerAccount) EarnLoyaltyPoints(points int) {
    c.LoyaltyPoints += points
}

func (c *CustomerAccount) RedeemLoyaltyPoints(points int) {
    c.LoyaltyPoints -= points
}

func (c *CustomerAccount) MintNFT(details string) {
    nft := NFT{ID: generateNFTID(), Details: details}
    c.NFTCollection = append(c.NFTCollection, nft)
}

func (c *CustomerAccount) TransferNFT(nftID, recipientID string) {
    // Implement NFT transfer logic
}

func (c *CustomerAccount) Save() error {
    log.Println("Saving customer:", c)
    if config.DB == nil {
        log.Println("Database connection is nil")
        return fmt.Errorf("database connection is nil")
    }
    _, err := config.DB.Exec(context.Background(), "INSERT INTO customers (customer_id, name, email, encrypted_payment_details, preloaded_balance, loyalty_points) VALUES ($1, $2, $3, $4, $5, $6)",
        c.CustomerID, c.Name, c.Email, c.EncryptedPaymentDetails, c.PreloadedBalance, c.LoyaltyPoints)
    if err != nil {
        log.Printf("Failed to save customer: %v\n", err)
        return err
    }
    log.Println("Customer saved successfully")
    return nil
}

func GetCustomerByID(customerID string) (*CustomerAccount, error) {
    var customer CustomerAccount
    err := config.DB.QueryRow(context.Background(),
        "SELECT customer_id, name, email, encrypted_payment_details, preloaded_balance, loyalty_points FROM customers WHERE customer_id=$1", customerID).
        Scan(&customer.CustomerID, &customer.Name, &customer.Email, &customer.EncryptedPaymentDetails, &customer.PreloadedBalance, &customer.LoyaltyPoints)
    return &customer, err
}
