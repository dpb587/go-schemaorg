package paymentchargespecification

import "github.com/dpb587/go-schemaorg"

var (
	// The payment method(s) to which the payment charge specification applies.
	AppliesToPaymentMethod = schemaorg.NewProperty("appliesToPaymentMethod")

	// The delivery method(s) to which the delivery charge or payment charge
	// specification applies.
	AppliesToDeliveryMethod = schemaorg.NewProperty("appliesToDeliveryMethod")
)
