package highschool

import "github.com/dpb587/go-schemaorg"

// // A high school.
var Type = schemaorg.NewDataType("http://schema.org", "HighSchool")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
