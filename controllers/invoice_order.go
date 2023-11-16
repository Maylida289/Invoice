package controllers

import (
	"Invoice/db"
	"Invoice/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type InvoiceReq struct { //untuk memparsing data yang diminta client sehingga dapat diproses ke sarvernya
	ID_customer      int     `json:"id_customer" param:"id"`
	Subject          string  `json:"subject" validate:"required"`
	Start_date       string  `json:"start_date" validate:"required"`
	Due_date         string  `json:"due_date" validate:"required"`
	Customer_name    string  `json:"customer_name" validate:"required"`
	Customer_address string  `json:"customer_address" validate:"required"`
	Item_name        string  `json:"item_name" validate:"required"`
	Quantity         float64 `json:"quantity" validate:"required"`
	Total_price      float64 `json:"total_price" validate:"required"`
}
type BayarReq struct {
	ID_bayar   int     `json:"id_bayar" param:"id"`
	TotalAmout float64 `json:"total" validate:"required"`
}

var CustomerList = make(map[int]models.Customer, 0)

var LastInvoice int = 1

func CreateCustomer(c echo.Context) (err error) { //menambahkan data konsumen
	req := new(InvoiceReq)
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
		// ID:        LastTotal,
		Subject:          req.Subject,
		Start_date:       req.Start_date,
		Due_date:         req.Due_date,
		Customer_name:    req.Customer_name,
		Customer_address: req.Customer_address,
		Item_name:        req.Item_name,
		Quantity:         req.Quantity,
		Total_price:      req.Total_price,
	}

	CustomerList[LastInvoice] = total
	LastInvoice++

	fmt.Println("user:", total)
	result := db.Create(&total)

	return c.JSON(http.StatusCreated, result)
}

func GetInvoice(c echo.Context) (err error) { //menampilkan seluruh data konsumen
	//static API
	db := db.DBManager()

	var invoices []models.Customer
	if err := db.Find(&invoices).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var responseInvoices []models.ResponseInvoice
	for _, invoice := range invoices {
		responseInvoice := models.ResponseInvoice{
			Subject:          invoice.Subject,
			Start_date:       invoice.Start_date,
			Due_date:         invoice.Due_date,
			Customer_name:    invoice.Customer_name,
			Customer_address: invoice.Customer_address,
			Total:            invoice.Total_price,
		}

		responseInvoices = append(responseInvoices, responseInvoice)
	}

	return c.JSON(http.StatusOK, responseInvoices)
}
