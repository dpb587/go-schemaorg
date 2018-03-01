package defenceestablishment

import "github.com/dpb587/go-schemaorg"

// // A defence establishment, such as an army or navy base.
var Type = schemaorg.NewDataType("http://schema.org", "DefenceEstablishment")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
