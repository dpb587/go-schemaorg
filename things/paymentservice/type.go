package paymentservice

import "github.com/dpb587/go-schemaorg"

// // A Service to transfer funds from a person or organization to a beneficiary
// person or organization.
var Type = schemaorg.NewDataType("http://schema.org", "PaymentService")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
