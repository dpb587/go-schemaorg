package course

import "github.com/dpb587/go-schemaorg"

// // A description of an educational course which may be offered as distinct
// instances at which take place at different times or take place at different
// locations, or be offered through different media or modes of study. An
// educational course is a sequence of one or more educational events and/or
// creative works which aims to build knowledge, competence or ability of
// learners.
var Type = schemaorg.NewDataType("http://schema.org", "Course")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
