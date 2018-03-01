package enginespecification

import "github.com/dpb587/go-schemaorg"

// // Information about the engine of the vehicle. A vehicle can have multiple
// engines represented by multiple engine specification entities.
var Type = schemaorg.NewDataType("http://schema.org", "EngineSpecification")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
