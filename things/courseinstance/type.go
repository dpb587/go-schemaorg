package courseinstance

import "github.com/dpb587/go-schemaorg"

// // An instance of a <a class="localLink"
// href="http://schema.org/Course">Course</a> which is distinct from other
// instances because it is offered at a different time or location or through
// different media or modes of study or to a specific section of students.
var Type = schemaorg.NewDataType("http://schema.org", "CourseInstance")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
