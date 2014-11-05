package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/thetruetrade/gotrade-db/models"
	"log"
	"fmt"
	"os"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	copyCommand := fmt.Sprintf("COPY currency(i_s_o3_code, i_s_o_numeric_code, name) FROM '%s\\data\\gotrade-currencies - gotrade-currencies.csv' CSV HEADER;", path)

	db, err := gorm.Open("postgres", "user=gotradeuser password=gotradeuser dbname=gotrade sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(true)

	// Disable table name's pluralization
	db.SingularTable(true)

	db.CreateTable(new(models.Currency))

	db.Exec(copyCommand)

//	db.CreateTable(new(models.Country))
//
//	country := models.Country{
//		ISO2Code: "AD",
//		ISO3Code: "AND",
//		ISONumericCode: 20,
//		Name: "Andorra",
//		TLDCode: ".ad",
//		Currency: models.Currency{ISO3Code: "EUR", ISONumericCode: 978, Name: "Euro"},
//	}
//	db.Save(&country)
//
//	country = models.Country{
//		ISO2Code: "AE",
//		ISO3Code: "ARE",
//		ISONumericCode: 784,
//		Name: "United Arab Emirates",
//		TLDCode: ".ae",
//		Currency: models.Currency{ISO3Code: "AED", ISONumericCode: 784, Name: "United Arab Emirates dirham"},
//	}
//	db.Save(&country)
//
//	country = models.Country{
//		ISO2Code: "AF",
//		ISO3Code: "AFG",
//		ISONumericCode: 4,
//		Name: "Afghanistan",
//		TLDCode: ".af",
//		Currency: models.Currency{ISO3Code: "AFN", ISONumericCode: 971, Name: "Afghani"},
//	}
//	db.Save(&country)
//
//	country = models.Country{
//		ISO2Code: "AG",
//		ISO3Code: "ATG",
//		ISONumericCode: 28,
//		Name: "Antigua and Barbuda",
//		TLDCode: ".ag",
//		Currency: models.Currency{ISO3Code: "XCD", ISONumericCode: 951, Name: "East Caribbean dollar"},
//	}
//	db.Save(&country)
//
//	country = models.Country{
//		ISO2Code: "AI",
//		ISO3Code: "AIA",
//		ISONumericCode: 660,
//		Name: "Anguilla",
//		TLDCode: ".ai",
//		Currency: models.Currency{ISO3Code: "XCD", ISONumericCode: 951, Name: "East Caribbean dollar"},
//	}
//	db.Save(&country)
//
//	country = models.Country{
//		ISO2Code: "ZA",
//		ISO3Code: "ZAF",
//		ISONumericCode: 710,
//		Name: "South Africa",
//		TLDCode: ".za",
//		Currency: models.Currency{ISO3Code: "ZAR", ISONumericCode: 710, Name: "South African rand"},
//	}
//	db.Save(&country)
}
