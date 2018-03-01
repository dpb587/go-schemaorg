package compoundpricespecification

import "github.com/dpb587/go-schemaorg"

// // A compound price specification is one that bundles multiple prices that all
// apply in combination for different dimensions of consumption. Use the name
// property of the attached unit price specification for indicating the
// dimension of a price component (e.g. "electricity" or "final cleaning").
var Type = schemaorg.NewDataType("http://schema.org", "CompoundPriceSpecification")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
