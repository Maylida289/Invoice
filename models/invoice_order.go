package models

type Invoice struct {
	ID_invoice       uint
	Subject          string
	Start_date       string
	Due_date         string
	Customer_name    string
	Customer_address string
	Item_name        string
	Quantity         float64
	Unit_Price       float64
	Total_price      float64
	Grand_total      float64
}
