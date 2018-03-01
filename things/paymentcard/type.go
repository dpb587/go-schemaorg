package paymentcard

import "github.com/dpb587/go-schemaorg"

// // A payment method using a credit, debit, store or other card to associate the
// payment with an account.
var Type = schemaorg.NewDataType("http://schema.org", "PaymentCard")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
