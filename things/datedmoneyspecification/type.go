package datedmoneyspecification

import "github.com/dpb587/go-schemaorg"

// // A DatedMoneySpecification represents monetary values with optional start and
// end dates. For example, this could represent an employee's salary over a
// specific period of time. <strong>Note:</strong> This type has been superseded
// by <a class="localLink"
// href="http://schema.org/MonetaryAmount">MonetaryAmount</a> use of that type
// is recommended
var Type = schemaorg.NewDataType("http://schema.org", "DatedMoneySpecification")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
