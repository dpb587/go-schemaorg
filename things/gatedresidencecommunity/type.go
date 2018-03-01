package gatedresidencecommunity

import "github.com/dpb587/go-schemaorg"

// // Residence type: Gated community.
var Type = schemaorg.NewDataType("http://schema.org", "GatedResidenceCommunity")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
