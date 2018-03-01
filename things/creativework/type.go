package creativework

import "github.com/dpb587/go-schemaorg"

// // The most generic kind of creative work, including books, movies, photographs,
// software programs, etc.
var Type = schemaorg.NewDataType("http://schema.org", "CreativeWork")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
