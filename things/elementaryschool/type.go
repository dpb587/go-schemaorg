package elementaryschool

import "github.com/dpb587/go-schemaorg"

// // An elementary school.
var Type = schemaorg.NewDataType("http://schema.org", "ElementarySchool")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
