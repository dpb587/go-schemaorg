package offeritemcondition

import "github.com/dpb587/go-schemaorg"

// // A list of possible conditions for the item.
var Type = schemaorg.NewDataType("http://schema.org", "OfferItemCondition")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
