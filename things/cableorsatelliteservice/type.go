package cableorsatelliteservice

import "github.com/dpb587/go-schemaorg"

// // A service which provides access to media programming like TV or radio. Access
// may be via cable or satellite.
var Type = schemaorg.NewDataType("http://schema.org", "CableOrSatelliteService")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
