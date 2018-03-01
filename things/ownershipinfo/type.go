package ownershipinfo

import "github.com/dpb587/go-schemaorg"

// // A structured value providing information about when a certain organization or
// person owned a certain product.
var Type = schemaorg.NewDataType("http://schema.org", "OwnershipInfo")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
