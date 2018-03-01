package drivewheelconfigurationvalue

import "github.com/dpb587/go-schemaorg"

// // A value indicating which roadwheels will receive torque.
var Type = schemaorg.NewDataType("http://schema.org", "DriveWheelConfigurationValue")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
