package receiveaction

import "github.com/dpb587/go-schemaorg"

var (
	// A sub property of instrument. The method of delivery.
	DeliveryMethod = schemaorg.NewProperty("deliveryMethod")

	// A sub property of participant. The participant who is at the sending end of
	// the action.
	Sender = schemaorg.NewProperty("sender")
)
