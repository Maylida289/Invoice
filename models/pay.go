package models

import "github.com/jinzhu/gorm"

type Bayar struct {
	gorm.Model
	ID_bayar   int
	CustomerID int
	Customer   Customer `gorm:"ForeignKey:ID_customer;references:CustomerID"`
	TotalAmout float64
}
