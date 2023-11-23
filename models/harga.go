package models

type Harga struct {
	ID           int
	HargaPerItem float64
	QuantityID   int
	Quantity     Quantity `gorm:"ForeignKey:ID;references:QuantityID"`
}
