package flightreservation

import "github.com/dpb587/go-schemaorg"

var (
	// The airline-specific indicator of boarding order / preference.
	BoardingGroup = schemaorg.NewProperty("boardingGroup")

	// The passenger's sequence number as assigned by the airline.
	PassengerSequenceNumber = schemaorg.NewProperty("passengerSequenceNumber")

	// The type of security screening the passenger is subject to.
	SecurityScreening = schemaorg.NewProperty("securityScreening")

	// The priority status assigned to a passenger for security or boarding (e.g.
	// FastTrack or Priority).
	PassengerPriorityStatus = schemaorg.NewProperty("passengerPriorityStatus")
)
