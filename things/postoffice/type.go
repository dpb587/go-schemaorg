package postoffice

import "github.com/dpb587/go-schemaorg"

// // A post office.
var Type = schemaorg.NewDataType("http://schema.org", "PostOffice")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
