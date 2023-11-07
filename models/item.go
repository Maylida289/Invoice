package models

type Item struct {
	ID_item        uint
	Item_name      string
	Quantity       float64
	Unit_Price     float64
	Total_price    float64
	Grand_total    float64
	InvoiceID      uint
	Invoice_detail Invoice_detail `gorm:"ForeignKey:ID;references:UserID"`
}
