package emergencyservice

import "github.com/dpb587/go-schemaorg"

// // An emergency service, such as a fire station or ER.
var Type = schemaorg.NewDataType("http://schema.org", "EmergencyService")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
