package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "project/handlers"
	"project/config"
)

func main() {

	// Initialize the database
	config.InitDB()

    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Routes
    e.GET("/", func(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]string{"message": "Welcome to the API"})
    })
    e.POST("/customers", handlers.CreateCustomer)
    e.GET("/customers/:id", handlers.GetCustomer)
    e.PUT("/customers/:id", handlers.UpdateCustomer)
    e.DELETE("/customers/:id", handlers.DeleteCustomer)

    // Start server
    e.Logger.Fatal(e.Start(":8080"))
}
