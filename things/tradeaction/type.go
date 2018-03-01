package tradeaction

import "github.com/dpb587/go-schemaorg"

// // The act of participating in an exchange of goods and services for monetary
// compensation. An agent trades an object, product or service with a
// participant in exchange for a one time or periodic payment.
var Type = schemaorg.NewDataType("http://schema.org", "TradeAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
