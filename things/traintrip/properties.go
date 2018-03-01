package traintrip

import "github.com/dpb587/go-schemaorg"

var (
	// The station from which the train departs.
	DepartureStation = schemaorg.NewProperty("departureStation")

	// The service provider, service operator, or service performer; the goods
	// producer. Another party (a seller) may offer those services or goods on
	// behalf of the provider. A provider may also serve as the seller.
	Provider = schemaorg.NewProperty("provider")

	// The platform where the train arrives.
	ArrivalPlatform = schemaorg.NewProperty("arrivalPlatform")

	// The platform from which the train departs.
	DeparturePlatform = schemaorg.NewProperty("departurePlatform")

	// The name of the train (e.g. The Orient Express).
	TrainName = schemaorg.NewProperty("trainName")

	// The unique identifier for the train.
	TrainNumber = schemaorg.NewProperty("trainNumber")

	// The station where the train trip ends.
	ArrivalStation = schemaorg.NewProperty("arrivalStation")

	// The expected arrival time.
	ArrivalTime = schemaorg.NewProperty("arrivalTime")

	// The expected departure time.
	DepartureTime = schemaorg.NewProperty("departureTime")
)
