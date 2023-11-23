package models

type Quantity struct {
	ID         int
	Quantity   float64
	CustomerID int
	Customer   Customer `gorm:"ForeignKey:ID;references:CustomerID"`
}
