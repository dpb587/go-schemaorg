package listenaction

import "github.com/dpb587/go-schemaorg"

// // The act of consuming audio content.
var Type = schemaorg.NewDataType("http://schema.org", "ListenAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
