package paymentchargespecification

import "github.com/dpb587/go-schemaorg"

// // The costs of settling the payment using a particular payment method.
var Type = schemaorg.NewDataType("http://schema.org", "PaymentChargeSpecification")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
