package offercatalog

import "github.com/dpb587/go-schemaorg"

// // An OfferCatalog is an ItemList that contains related Offers and/or further
// OfferCatalogs that are offeredBy the same provider.
var Type = schemaorg.NewDataType("http://schema.org", "OfferCatalog")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
