package currencyconversionservice

import "github.com/dpb587/go-schemaorg"

// // A service to convert funds from one currency to another currency.
var Type = schemaorg.NewDataType("http://schema.org", "CurrencyConversionService")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
