package controllers

import (
	database "Invoice/db"
	"Invoice/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type InvoiceReq struct { //untuk memparsing data yang diminta client sehingga dapat diproses ke sarvernya
	ID_invoice uint   `json:"id" param:"id"`
	Subject    string `json:"suject" validate:"required"`
	Start_date string `json:"start_date" validate:"required"`
	Due_date   string `json:"due_date" validate:"required"`
}
type CustomerReq struct { //untuk memparsing data yang diminta client sehingga dapat diproses ke sarvernya
	ID_customer      uint   `json:"id" param:"id"`
	Customer_name    string `json:"customer_name" validate:"required"`
	Customer_address string `json:"customer_address" validate:"required"`
}
type itemsReq struct { //untuk memparsing data yang diminta client sehingga dapat diproses ke sarvernya
	ID_item     uint    `json:"id" param:"id"`
	Item_name   float64 `json:"item_name" validate:"required"`
	Quantity    float64 `json:"quantity" validate:"required"`
	Unit_Price  float64 `json:"unit_price" validate:"required"`
	Total_price float64 `json:"total_price" validate:"required"`
	Grand_total float64 `json:"grand_total" validate:"required"`
}

var InvoiceList = make(map[int]models.Invoice_detail, 0)
var CustomerList = make(map[int]models.Customer, 0)
var ItemList = make(map[int]models.Item, 0)

var LastInvoice int = 1
var LastCustomer int = 1
var LastItem int = 1

func CreateInvoice(c echo.Context) (err error) { //menambahkan data konsumen
	req := new(InvoiceReq)
	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if err = c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	//static API
	invoice := models.Invoice_detail{
		ID_invoice: uint(LastInvoice),
		Subject:    req.Subject,
		Start_date: req.Start_date,
		Due_date:   req.Due_date,
	}
	InvoiceList[LastInvoice] = invoice
	LastInvoice++

	db := database.DBManager()
	// if err := db.Create(&admin).Error; err != nil {
	// 	// Handle kesalahan yang terjadi saat menyimpan ke database
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }
	result := db.Create(&invoice)

	return c.JSON(http.StatusCreated, result)
}

func CreateCustomer(c echo.Context) (err error) { //menambahkan data konsumen
	req := new(CustomerReq)
	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if err = c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	//static API
	customer := models.Customer{
		ID_customer:      uint(LastCustomer),
		Customer_name:    req.Customer_name,
		Customer_address: req.Customer_address,
	}
	CustomerList[LastCustomer] = customer
	LastCustomer++

	db := database.DBManager()
	if err := db.Create(&customer).Error; err != nil {
		// Handle kesalahan yang terjadi saat menyimpan ke database
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	result := db.Create(&customer)

	return c.JSON(http.StatusCreated, result)
}
