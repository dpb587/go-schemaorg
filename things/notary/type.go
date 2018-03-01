package notary

import "github.com/dpb587/go-schemaorg"

// // A notary.
var Type = schemaorg.NewDataType("http://schema.org", "Notary")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
