package datatype

import "github.com/dpb587/go-schemaorg"

// // The basic data types such as Integers, Strings, etc.
var Type = schemaorg.NewDataType("http://schema.org", "DataType")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
