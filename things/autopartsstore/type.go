package autopartsstore

import "github.com/dpb587/go-schemaorg"

// // An auto parts store.
var Type = schemaorg.NewDataType("http://schema.org", "AutoPartsStore")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
