package bustrip

import "github.com/dpb587/go-schemaorg"

var (
	// The service provider, service operator, or service performer; the goods
	// producer. Another party (a seller) may offer those services or goods on
	// behalf of the provider. A provider may also serve as the seller.
	Provider = schemaorg.NewProperty("provider")

	// The name of the bus (e.g. Bolt Express).
	BusName = schemaorg.NewProperty("busName")

	// The stop or station from which the bus departs.
	DepartureBusStop = schemaorg.NewProperty("departureBusStop")

	// The stop or station from which the bus arrives.
	ArrivalBusStop = schemaorg.NewProperty("arrivalBusStop")

	// The expected arrival time.
	ArrivalTime = schemaorg.NewProperty("arrivalTime")

	// The unique identifier for the bus.
	BusNumber = schemaorg.NewProperty("busNumber")

	// The expected departure time.
	DepartureTime = schemaorg.NewProperty("departureTime")
)
