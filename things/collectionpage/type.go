package collectionpage

import "github.com/dpb587/go-schemaorg"

// // Web page type: Collection page.
var Type = schemaorg.NewDataType("http://schema.org", "CollectionPage")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
