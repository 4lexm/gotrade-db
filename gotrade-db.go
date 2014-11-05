package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/thetruetrade/gotrade-db/models"
	"log"
)

func main() {
	db, err := gorm.Open("postgres", "user=gotradeuser password=gotradeuser dbname=gotrade sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(true)

	// Disable table name's pluralization
	db.SingularTable(true)

	db.CreateTable(new(models.Currency))
	Currency1 := models.Currency{ISO3Code: "ZAR", ISONumericCode: 710, Name: "South African rand"}
	db.Save(&Currency1)

	//	var zar Currency
	//	db.Where("name = ?", "South African rand").Find(&zar)

	//	db.CreateTable(new(Country))
	//	Country1 := Country{ISO2Code: "ZA", ISO3Code: "ZAF", ISONumericCode: 710, Name: "South Africa", TLDCode: ".za", Currency}

}
