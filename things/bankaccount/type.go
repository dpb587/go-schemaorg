package bankaccount

import "github.com/dpb587/go-schemaorg"

// // A product or service offered by a bank whereby one may deposit, withdraw or
// transfer money and in some cases be paid interest.
var Type = schemaorg.NewDataType("http://schema.org", "BankAccount")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
