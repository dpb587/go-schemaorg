package pricespecification

import "github.com/dpb587/go-schemaorg"

// // A structured value representing a price or price range. Typically, only the
// subclasses of this type are used for markup. It is recommended to use <a
// class="localLink" href="http://schema.org/MonetaryAmount">MonetaryAmount</a>
// to describe independent amounts of money such as a salary, credit card
// limits, etc.
var Type = schemaorg.NewDataType("http://schema.org", "PriceSpecification")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
