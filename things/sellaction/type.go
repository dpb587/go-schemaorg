package sellaction

import "github.com/dpb587/go-schemaorg"

// // The act of taking money from a buyer in exchange for goods or services
// rendered. An agent sells an object, product, or service to a buyer for a
// price. Reciprocal of BuyAction.
var Type = schemaorg.NewDataType("http://schema.org", "SellAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
