package models

type Country struct {
	Id 				int
	ISO2Code 		string `sql:"size:2"`
	ISO3Code 		string `sql:"size:3"`
	ISONumericCode 	int
	Name 			string `sql:"size:255"`
	TLDCode 		string `sql:"size:5"`
	Currency 		Currency
	CurrencyId 		int
}
