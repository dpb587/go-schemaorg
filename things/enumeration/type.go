package enumeration

import "github.com/dpb587/go-schemaorg"

// // Lists or enumerations—for example, a list of cuisines or music genres, etc.
var Type = schemaorg.NewDataType("http://schema.org", "Enumeration")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
