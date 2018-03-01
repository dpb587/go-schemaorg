package demand

import "github.com/dpb587/go-schemaorg"

// // A demand entity represents the public, not necessarily binding, not
// necessarily exclusive, announcement by an organization or person to seek a
// certain type of goods or services. For describing demand using this type, the
// very same properties used for Offer apply.
var Type = schemaorg.NewDataType("http://schema.org", "Demand")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
