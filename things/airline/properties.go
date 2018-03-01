package airline

import "github.com/dpb587/go-schemaorg"

var (
	// IATA identifier for an airline or airport.
	IataCode = schemaorg.NewProperty("iataCode")

	// The type of boarding policy used by the airline (e.g. zone-based or
	// group-based).
	BoardingPolicy = schemaorg.NewProperty("boardingPolicy")
)
