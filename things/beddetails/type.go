package beddetails

import "github.com/dpb587/go-schemaorg"

// // An entity holding detailed information about the available bed types, e.g.
// the quantity of twin beds for a hotel room. For the single case of just one
// bed of a certain type, you can use bed directly with a text. See also <a
// class="localLink" href="http://schema.org/BedType">BedType</a> (under
// development).
var Type = schemaorg.NewDataType("http://schema.org", "BedDetails")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
