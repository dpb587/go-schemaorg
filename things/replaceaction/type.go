package replaceaction

import "github.com/dpb587/go-schemaorg"

// // The act of editing a recipient by replacing an old object with a new object.
var Type = schemaorg.NewDataType("http://schema.org", "ReplaceAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
