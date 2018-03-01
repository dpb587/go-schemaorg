package transferaction

import "github.com/dpb587/go-schemaorg"

// // The act of transferring/moving (abstract or concrete) animate or inanimate
// objects from one place to another.
var Type = schemaorg.NewDataType("http://schema.org", "TransferAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
