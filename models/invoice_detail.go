package models

type Invoice_detail struct {
	ID_invoice uint
	Subject    string
	Start_date string
	Due_date   string
	CustomerID uint
	Customer   Customer `gorm:"ForeignKey:ID_customer;references:CustomerID"`
}
