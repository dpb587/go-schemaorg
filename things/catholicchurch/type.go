package catholicchurch

import "github.com/dpb587/go-schemaorg"

// // A Catholic church.
var Type = schemaorg.NewDataType("http://schema.org", "CatholicChurch")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
