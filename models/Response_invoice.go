package models

type ResponseInvoice struct {
	Subject          string
	Start_date       string
	Due_date         string
	Customer_name    string
	Customer_address string
	Total            float64
}
