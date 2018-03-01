package middleschool

import "github.com/dpb587/go-schemaorg"

// // A middle school (typically for children aged around 11-14, although this
// varies somewhat).
var Type = schemaorg.NewDataType("http://schema.org", "MiddleSchool")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
