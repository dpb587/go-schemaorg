package boardingpolicytype

import "github.com/dpb587/go-schemaorg"

// // A type of boarding policy used by an airline.
var Type = schemaorg.NewDataType("http://schema.org", "BoardingPolicyType")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
