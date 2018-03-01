package tipaction

import "github.com/dpb587/go-schemaorg"

// // The act of giving money voluntarily to a beneficiary in recognition of
// services rendered.
var Type = schemaorg.NewDataType("http://schema.org", "TipAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
