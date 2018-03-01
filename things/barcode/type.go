package barcode

import "github.com/dpb587/go-schemaorg"

// // An image of a visual machine-readable code such as a barcode or QR code.
var Type = schemaorg.NewDataType("http://schema.org", "Barcode")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
