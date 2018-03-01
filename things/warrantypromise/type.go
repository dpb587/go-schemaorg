package warrantypromise

import "github.com/dpb587/go-schemaorg"

// // A structured value representing the duration and scope of services that will
// be provided to a customer free of charge in case of a defect or malfunction
// of a product.
var Type = schemaorg.NewDataType("http://schema.org", "WarrantyPromise")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
