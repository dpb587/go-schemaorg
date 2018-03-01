package performingartstheater

import "github.com/dpb587/go-schemaorg"

// // A theater or other performing art center.
var Type = schemaorg.NewDataType("http://schema.org", "PerformingArtsTheater")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
