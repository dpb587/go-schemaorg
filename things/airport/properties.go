package airport

import "github.com/dpb587/go-schemaorg"

var (
	// IATA identifier for an airline or airport.
	IataCode = schemaorg.NewProperty("iataCode")

	// ICAO identifier for an airport.
	IcaoCode = schemaorg.NewProperty("icaoCode")
)
