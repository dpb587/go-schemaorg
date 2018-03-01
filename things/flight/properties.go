package flight

import "github.com/dpb587/go-schemaorg"

var (
	// The airport where the flight terminates.
	ArrivalAirport = schemaorg.NewProperty("arrivalAirport")

	// The service provider, service operator, or service performer; the goods
	// producer. Another party (a seller) may offer those services or goods on
	// behalf of the provider. A provider may also serve as the seller.
	Provider = schemaorg.NewProperty("provider")

	// Identifier of the flight's arrival gate.
	ArrivalGate = schemaorg.NewProperty("arrivalGate")

	// The unique identifier for a flight including the airline IATA code. For
	// example, if describing United flight 110, where the IATA code for United is
	// 'UA', the flightNumber is 'UA110'.
	FlightNumber = schemaorg.NewProperty("flightNumber")

	// 'carrier' is an out-dated term indicating the 'provider' for parcel delivery
	// and flights.
	Carrier = schemaorg.NewProperty("carrier")

	// The type of boarding policy used by the airline (e.g. zone-based or
	// group-based).
	BoardingPolicy = schemaorg.NewProperty("boardingPolicy")

	// The kind of aircraft (e.g., "Boeing 747").
	Aircraft = schemaorg.NewProperty("aircraft")

	// The expected arrival time.
	ArrivalTime = schemaorg.NewProperty("arrivalTime")

	// An entity which offers (sells / leases / lends / loans) the services / goods.
	// A seller may also be a provider.
	Seller = schemaorg.NewProperty("seller")

	// The time when a passenger can check into the flight online.
	WebCheckinTime = schemaorg.NewProperty("webCheckinTime")

	// Description of the meals that will be provided or available for purchase.
	MealService = schemaorg.NewProperty("mealService")

	// Identifier of the flight's departure gate.
	DepartureGate = schemaorg.NewProperty("departureGate")

	// Identifier of the flight's departure terminal.
	DepartureTerminal = schemaorg.NewProperty("departureTerminal")

	// The estimated time the flight will take.
	EstimatedFlightDuration = schemaorg.NewProperty("estimatedFlightDuration")

	// The distance of the flight.
	FlightDistance = schemaorg.NewProperty("flightDistance")

	// Identifier of the flight's arrival terminal.
	ArrivalTerminal = schemaorg.NewProperty("arrivalTerminal")

	// The airport where the flight originates.
	DepartureAirport = schemaorg.NewProperty("departureAirport")

	// The expected departure time.
	DepartureTime = schemaorg.NewProperty("departureTime")
)
