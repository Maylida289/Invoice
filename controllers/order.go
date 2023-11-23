package controllers

import (
	"Invoice/db"
	"Invoice/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HargaReq struct {
	ID           int     `json:"id_harga" gorm:"primaryKey" param:"id" validate:"required"`
	QuantityID   int     `json:"id_customer" validate:"required"`
	HargaPerItem float64 `json:"hargaper_item" validate:"required"`
}

type PelangganReq struct { //untuk memparsing data yang diminta client sehingga dapat diproses ke sarvernya
	ID               int    `json:"id_customer" gorm:"primaryKey" param:"id" validate:"required"`
	Subject          string `json:"subject"`
	Start_date       string `json:"start_date"`
	Due_date         string `json:"due_date"`
	Customer_name    string `json:"customer_name"`
	Customer_address string `json:"customer_address"`
}

type QuantityReq struct {
	ID         int     `json:"id_quantity" gorm:"primaryKey" param:"id" validate:"required"`
	CustomerID int     `json:"id_customer" validate:"required"`
	Quantity   float64 `json:"quantity" validate:"required"`
}

var CustomerList = make(map[int]models.Customer, 0)
var LastCustomer int = 1

var QuantityList = make(map[int]models.Quantity, 0)
var LastQuantity int = 1

var HargaList = make(map[int]models.Harga, 0)
var LastHarga int = 1

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

func CreateHarga(c echo.Context) (err error) {
	req := new(HargaReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db := db.DBManager()

	// Buat Harga untuk Quantity terkait dengan Customer
	harga := models.Harga{
		ID:           LastHarga,
		QuantityID:   req.QuantityID,
		HargaPerItem: req.HargaPerItem,
	}
	HargaList[LastHarga] = harga
	LastHarga++
	fmt.Println("harga:", harga)

	result := db.Create(&harga)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}

	return c.JSON(http.StatusCreated, harga)

	// if err := db.Create(&harga).Error; err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err.Error()) // aku mau bikin pesan di dalam sini supaya keliatan output di postman
	// } //pesannya nanti berupa response atau hasil dari outputnya nanti bikin 1 func response untuk melihat response nya seperti apa.

	// return c.JSON(http.StatusCreated, "Harga berhasil dibuat untuk Quantity terkait")
}

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
