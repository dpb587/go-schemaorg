package realestateagent

import "github.com/dpb587/go-schemaorg"

// // A real-estate agent.
var Type = schemaorg.NewDataType("http://schema.org", "RealEstateAgent")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
