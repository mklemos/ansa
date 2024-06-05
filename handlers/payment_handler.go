package handlers

import (
    "net/http"
    "project/models"
    "github.com/labstack/echo/v4"
)

func ProcessPayment(c echo.Context) error {
    paymentRequest := new(models.PaymentRequest)
    if err := c.Bind(paymentRequest); err != nil {
        return err
    }

    customer, err := models.GetCustomerByID(paymentRequest.CustomerID)
    if err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"message": "Customer not found"})
    }

    paymentProcessor := models.PaymentProcessing{}
    result := paymentProcessor.CreateChargeRequest(customer, paymentRequest.Amount)

    response := models.PaymentResponse{
        Message: result,
    }
    return c.JSON(http.StatusOK, response)
}
