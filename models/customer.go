package models

type Customer struct {
	ID_customer      uint
	Customer_name    string
	Customer_address string
	Invoice_detail   []Invoice_detail
}
