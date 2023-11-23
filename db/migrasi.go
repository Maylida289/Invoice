package db

import (
	"Invoice/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Config struct {
	Username     string
	Password     string
	Port         string
	Host         string
	DatabaseName string
}

var db *gorm.DB
var err error

func Init() {
	config := Config{
		Username:     "root",
		Password:     "",
		Host:         "tcp(127.0.0.1:3306)",
		DatabaseName: "hargain",
	}

	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", config.Username, config.Password, config.Host, config.DatabaseName))

	if err != nil {
		panic("database error")
	}
	db.AutoMigrate(&models.Customer{}, models.Harga{}, models.Quantity{}) //migrate schema

	db.Delete(&Customers)
	for key, _ := range Customers {
		db.Create(&Customers[key])
	}

	db.Delete(&Quantitys)
	for key, _ := range Quantitys {
		db.Create(&Quantitys[key])
	}

	db.Delete(&Hargas)
	for key, _ := range Hargas {
		db.Create(&Hargas[key])
	}

}
func DBManager() *gorm.DB {
	return db
}
