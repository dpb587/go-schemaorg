package financialservice

import "github.com/dpb587/go-schemaorg"

// // Financial services business.
var Type = schemaorg.NewDataType("http://schema.org", "FinancialService")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
