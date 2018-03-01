package paymentstatustype

import "github.com/dpb587/go-schemaorg"

// // A specific payment status. For example, PaymentDue, PaymentComplete, etc.
var Type = schemaorg.NewDataType("http://schema.org", "PaymentStatusType")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
