package restricteddiet

import "github.com/dpb587/go-schemaorg"

// // A diet restricted to certain foods or preparations for cultural, religious,
// health or lifestyle reasons.
var Type = schemaorg.NewDataType("http://schema.org", "RestrictedDiet")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
