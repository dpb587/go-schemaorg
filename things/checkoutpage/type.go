package checkoutpage

import "github.com/dpb587/go-schemaorg"

// // Web page type: Checkout page.
var Type = schemaorg.NewDataType("http://schema.org", "CheckoutPage")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
