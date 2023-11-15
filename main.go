package main

import (
	"Invoice/controllers"
	"Invoice/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	db.Init()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		result := map[string]string{
			"response": "200",
			"Messege":  "Success",
		}
		return c.JSON(http.StatusOK, result)
	})
	e.POST("/add/invoice", controllers.CreateInvoice)

	// Rute untuk membuat rincian faktur
	e.GET("/invoice_details", controllers.GetInvoice)

	// Rute untuk membuat item faktur
	// e.POST("/items", controllers.CreateItem)
	e.Logger.Fatal(e.Start(":8070"))
}
