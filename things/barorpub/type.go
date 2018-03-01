package barorpub

import "github.com/dpb587/go-schemaorg"

// // A bar or pub.
var Type = schemaorg.NewDataType("http://schema.org", "BarOrPub")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
