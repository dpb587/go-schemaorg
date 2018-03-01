package selfstorage

import "github.com/dpb587/go-schemaorg"

// // A self-storage facility.
var Type = schemaorg.NewDataType("http://schema.org", "SelfStorage")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
