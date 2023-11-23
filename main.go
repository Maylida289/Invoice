package main

import (
	"Invoice/controllers"
	"Invoice/db"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func main() {
	db.Init()
	e := echo.New()
	e.Validator = &CustomValidator{
		Validator: validator.New(),
	}

	e.GET("/", func(c echo.Context) error {
		result := map[string]string{
			"response": "200",
			"Messege":  "Success",
		}
		return c.JSON(http.StatusOK, result)
	})
	e.POST("/add/customer", controllers.CreateCustomer)

	//memasukan harga
	e.POST("/input/harga", controllers.CreateHarga)

	//memasukan harga dari masing-masing quantity
	e.POST("/input/quantity", controllers.CreateQuantity)

	// Rute untuk membuat item faktur
	// e.POST("/items", controllers.CreateItem)
	e.Logger.Fatal(e.Start("localhost:8076"))
}
