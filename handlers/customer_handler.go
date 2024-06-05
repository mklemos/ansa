package handlers

import (
    "context"
    "net/http"
    "project/config"
    "project/models"
    "github.com/labstack/echo/v4"
	//"fmt"
)

func CreateCustomer(c echo.Context) error {
    customer := new(models.CustomerAccount)
    if err := c.Bind(customer); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
    }
    if err := customer.Save(); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create customer", "error": err.Error()})
    }
    return c.JSON(http.StatusCreated, customer)
}

func GetCustomer(c echo.Context) error {
    id := c.Param("id")
    customer, err := models.GetCustomerByID(id)
    if err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"message": "Customer not found"})
    }
    return c.JSON(http.StatusOK, customer)
}

func UpdateCustomer(c echo.Context) error {
    id := c.Param("id")
    customer, err := models.GetCustomerByID(id)
    if err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"message": "Customer not found"})
    }
    if err := c.Bind(customer); err != nil {
        return err
    }
    if err := customer.Save(); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update customer"})
    }
    return c.JSON(http.StatusOK, customer)
}

func DeleteCustomer(c echo.Context) error {
    id := c.Param("id")
    _, err := models.GetCustomerByID(id)
    if err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"message": "Customer not found"})
    }
    _, err = config.DB.Exec(context.Background(), "DELETE FROM customers WHERE customer_id=$1", id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to delete customer"})
    }
    return c.NoContent(http.StatusNoContent)
}
