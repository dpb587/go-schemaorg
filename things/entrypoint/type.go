package entrypoint

import "github.com/dpb587/go-schemaorg"

// // An entry point, within some Web-based protocol.
var Type = schemaorg.NewDataType("http://schema.org", "EntryPoint")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
