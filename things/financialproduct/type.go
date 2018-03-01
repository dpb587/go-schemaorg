package financialproduct

import "github.com/dpb587/go-schemaorg"

// // A product provided to consumers and businesses by financial institutions such
// as banks, insurance companies, brokerage firms, consumer finance companies,
// and investment companies which comprise the financial services industry.
var Type = schemaorg.NewDataType("http://schema.org", "FinancialProduct")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
