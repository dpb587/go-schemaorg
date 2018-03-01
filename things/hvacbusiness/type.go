package hvacbusiness

import "github.com/dpb587/go-schemaorg"

// // A business that provide Heating, Ventilation and Air Conditioning services.
var Type = schemaorg.NewDataType("http://schema.org", "HVACBusiness")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
