package buyaction

import "github.com/dpb587/go-schemaorg"

// // The act of giving money to a seller in exchange for goods or services
// rendered. An agent buys an object, product, or service from a seller for a
// price. Reciprocal of SellAction.
var Type = schemaorg.NewDataType("http://schema.org", "BuyAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
