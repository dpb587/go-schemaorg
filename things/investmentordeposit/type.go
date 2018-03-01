package investmentordeposit

import "github.com/dpb587/go-schemaorg"

// // A type of financial product that typically requires the client to transfer
// funds to a financial service in return for potential beneficial financial
// return.
var Type = schemaorg.NewDataType("http://schema.org", "InvestmentOrDeposit")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
