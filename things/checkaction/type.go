package checkaction

import "github.com/dpb587/go-schemaorg"

// // An agent inspects, determines, investigates, inquires, or examines an
// object's accuracy, quality, condition, or state.
var Type = schemaorg.NewDataType("http://schema.org", "CheckAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
