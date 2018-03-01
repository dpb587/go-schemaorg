package drycleaningorlaundry

import "github.com/dpb587/go-schemaorg"

// // A dry-cleaning business.
var Type = schemaorg.NewDataType("http://schema.org", "DryCleaningOrLaundry")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
