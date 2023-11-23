package controllers

import (
	"Invoice/db"
	"Invoice/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PelangganReq struct { //untuk memparsing data yang diminta client sehingga dapat diproses ke sarvernya
	ID               int    `json:"id_customer" gorm:"primaryKey" param:"id" validate:"required"`
	Subject          string `json:"subject"`
	Start_date       string `json:"start_date"`
	Due_date         string `json:"due_date"`
	Customer_name    string `json:"customer_name"`
	Customer_address string `json:"customer_address"`
}

var CustomerList = make(map[int]models.Customer, 0)
var LastCustomer int = 1

func CreateCustomer(c echo.Context) (err error) { //menambahkan data konsumen
	req := new(PelangganReq)
	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if err = c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	//static API
	// admin := models.Total_price{
	// 	ID:        LastTotal,
	// 	Price:     req.Price,
	// 	Exchange:  req.Exchange,
	// 	Pair:      req.Pair,
	// 	PairPrice: req.PairPrice,
	// 	Volume:    req.Volume,
	// }
	// TotalList[LastTotal] = admin
	// LastTotal++

	db := db.DBManager()
	total := models.Customer{
		ID:               LastCustomer,
		Subject:          req.Subject,
		Start_date:       req.Start_date,
		Due_date:         req.Due_date,
		Customer_name:    req.Customer_name,
		Customer_address: req.Customer_address,
	}

	CustomerList[LastCustomer] = total
	LastCustomer++
	fmt.Println("user:", total)
	result := db.Create(&total)

	return c.JSON(http.StatusCreated, result)
}
