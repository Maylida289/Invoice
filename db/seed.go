package db

import "Invoice/models"

var Customers = []models.Customer{
	{
		ID:               1,
		Subject:          "Invoice Kain",
		Start_date:       "23 Mei 2022",
		Due_date:         "12 Juni 2022",
		Customer_name:    "Alpiansah",
		Customer_address: "Jakarta Barat",
	},
	{
		ID:               2,
		Subject:          "Invoice Perkakas",
		Start_date:       "20 September 2022",
		Due_date:         "30 September 2022",
		Customer_name:    "Derina Aldiri S.E",
		Customer_address: "Surakarta",
	},
}

var Quantitys = []models.Quantity{
	{
		ID:         1,
		Quantity:   5,
		CustomerID: 1,
	},
	{
		ID:         2,
		Quantity:   7,
		CustomerID: 1,
	},
}

var Hargas = []models.Harga{
	{
		ID:           1,
		HargaPerItem: 7000,
		QuantityID:   1,
	},
	{
		ID:           2,
		HargaPerItem: 10000,
		QuantityID:   2,
	},
}
