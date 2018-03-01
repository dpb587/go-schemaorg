package deleteaction

import "github.com/dpb587/go-schemaorg"

// // The act of editing a recipient by removing one of its objects.
var Type = schemaorg.NewDataType("http://schema.org", "DeleteAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
