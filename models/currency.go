package models

type Currency struct {
	Id             int
	ISO3Code       string `sql:"size:3"`
	ISONumericCode int
	Name           string `sql:"size:255"`
}
