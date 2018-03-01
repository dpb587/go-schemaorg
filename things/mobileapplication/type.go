package mobileapplication

import "github.com/dpb587/go-schemaorg"

// // A software application designed specifically to work well on a mobile device
// such as a telephone.
var Type = schemaorg.NewDataType("http://schema.org", "MobileApplication")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
