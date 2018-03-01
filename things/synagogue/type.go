package synagogue

import "github.com/dpb587/go-schemaorg"

// // A synagogue.
var Type = schemaorg.NewDataType("http://schema.org", "Synagogue")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
