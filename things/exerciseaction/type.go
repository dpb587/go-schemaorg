package exerciseaction

import "github.com/dpb587/go-schemaorg"

// // The act of participating in exertive activity for the purposes of improving
// health and fitness.
var Type = schemaorg.NewDataType("http://schema.org", "ExerciseAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
