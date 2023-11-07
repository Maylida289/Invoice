package main

import (
	"Invoice/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		result := map[string]string{
			"response": "200",
			"Messege":  "Success",
		}
		return c.JSON(http.StatusOK, result)
	})
	e.POST("/customers", controllers.CreateCustomer)

	// Rute untuk membuat rincian faktur
	e.POST("/invoice_details", controllers.CreateInvoice)

	// Rute untuk membuat item faktur
	// e.POST("/items", controllers.CreateItem)
	e.Logger.Fatal(e.Start(":8062"))
}
