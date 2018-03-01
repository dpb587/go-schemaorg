package postaladdress

import "github.com/dpb587/go-schemaorg"

// // The mailing address.
var Type = schemaorg.NewDataType("http://schema.org", "PostalAddress")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
