package sportinggoodsstore

import "github.com/dpb587/go-schemaorg"

// // A sporting goods store.
var Type = schemaorg.NewDataType("http://schema.org", "SportingGoodsStore")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
