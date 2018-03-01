package exercisegym

import "github.com/dpb587/go-schemaorg"

// // A gym.
var Type = schemaorg.NewDataType("http://schema.org", "ExerciseGym")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
