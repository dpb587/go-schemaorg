package sendaction

import "github.com/dpb587/go-schemaorg"

var (
	// A sub property of participant. The participant who is at the receiving end of
	// the action.
	Recipient = schemaorg.NewProperty("recipient")

	// A sub property of instrument. The method of delivery.
	DeliveryMethod = schemaorg.NewProperty("deliveryMethod")
)
