package qualitativevalue

import "github.com/dpb587/go-schemaorg"

// // A predefined value for a product characteristic, e.g. the power cord plug
// type 'US' or the garment sizes 'S', 'M', 'L', and 'XL'.
var Type = schemaorg.NewDataType("http://schema.org", "QualitativeValue")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
