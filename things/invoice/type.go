package invoice

import "github.com/dpb587/go-schemaorg"

// // A statement of the money due for goods or services; a bill.
var Type = schemaorg.NewDataType("http://schema.org", "Invoice")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
