package bankorcreditunion

import "github.com/dpb587/go-schemaorg"

// // Bank or credit union.
var Type = schemaorg.NewDataType("http://schema.org", "BankOrCreditUnion")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
