package rentaction

import "github.com/dpb587/go-schemaorg"

// // The act of giving money in return for temporary use, but not ownership, of an
// object such as a vehicle or property. For example, an agent rents a property
// from a landlord in exchange for a periodic payment.
var Type = schemaorg.NewDataType("http://schema.org", "RentAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
