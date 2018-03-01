package datacatalog

import "github.com/dpb587/go-schemaorg"

// // A collection of datasets.
var Type = schemaorg.NewDataType("http://schema.org", "DataCatalog")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
