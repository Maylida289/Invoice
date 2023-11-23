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

var HargaList = make(map[int]models.Harga, 0)
var LastHarga int = 1

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
