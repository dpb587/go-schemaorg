package internetcafe

import "github.com/dpb587/go-schemaorg"

// // An internet cafe.
var Type = schemaorg.NewDataType("http://schema.org", "InternetCafe")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
