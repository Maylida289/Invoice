package controllers

import (
	"Invoice/db"
	"Invoice/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type QuantityReq struct {
	ID         int     `json:"id_quantity" gorm:"primaryKey" param:"id" validate:"required"`
	CustomerID int     `json:"id_customer" validate:"required"`
	Quantity   float64 `json:"quantity" validate:"required"`
}

var QuantityList = make(map[int]models.Quantity, 0)
var LastQuantity int = 1

func CreateQuantity(c echo.Context) (err error) {
	req := new(QuantityReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db := db.DBManager()

	quantity := models.Quantity{
		ID:         LastQuantity,
		CustomerID: req.CustomerID,
		Quantity:   req.Quantity,
	}
	QuantityList[LastQuantity] = quantity
	LastQuantity++
	fmt.Println("quantity:", quantity)

	result := db.Create(&quantity)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}

	return c.JSON(http.StatusCreated, quantity)
}
