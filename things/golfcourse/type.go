package golfcourse

import "github.com/dpb587/go-schemaorg"

// // A golf course.
var Type = schemaorg.NewDataType("http://schema.org", "GolfCourse")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
