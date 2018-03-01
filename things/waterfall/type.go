package waterfall

import "github.com/dpb587/go-schemaorg"

// // A waterfall, like Niagara.
var Type = schemaorg.NewDataType("http://schema.org", "Waterfall")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
