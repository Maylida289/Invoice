package models

type Item struct { //ID Customer ada di item
	ID_item     int
	Item_name   string
	Quantity    float64
	Unit_Price  float64
	Total_price float64
	Grand_total float64
	CustomerID  int
	Customer    Customer `gorm:"ForeignKey:ID;references:CustomerID"`
}
