package governmentpermit

import "github.com/dpb587/go-schemaorg"

// // A permit issued by a government agency.
var Type = schemaorg.NewDataType("http://schema.org", "GovernmentPermit")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
