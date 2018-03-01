package installaction

import "github.com/dpb587/go-schemaorg"

// // The act of installing an application.
var Type = schemaorg.NewDataType("http://schema.org", "InstallAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
