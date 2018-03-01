package table

import "github.com/dpb587/go-schemaorg"

// // A table on a Web page.
var Type = schemaorg.NewDataType("http://schema.org", "Table")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
