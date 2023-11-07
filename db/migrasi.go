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
		DatabaseName: "invoice",
	}

	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", config.Username, config.Password, config.Host, config.DatabaseName))

	if err != nil {
		panic("database error")
	}
	db.AutoMigrate(&models.Customer{}, models.Invoice_detail{}, models.Item{}) //migrate schema

	// db.Delete(&Customer)
	// for key, _ := range Customer {
	// 	db.Create(&Customer[key])
	// }
	// db.Delete(&Invoice_detail)
	// for key, _ := range Invoice_detail {
	// 	db.Create(&Invoice_detail[key])
	// }
	// db.Delete(&Item)
	// for key, _ := range Item {
	// 	db.Create(&Item[key])
	// }
}
func DBManager() *gorm.DB {
	return db
}
