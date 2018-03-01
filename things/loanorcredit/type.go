package loanorcredit

import "github.com/dpb587/go-schemaorg"

// // A financial product for the loaning of an amount of money under agreed terms
// and charges.
var Type = schemaorg.NewDataType("http://schema.org", "LoanOrCredit")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
