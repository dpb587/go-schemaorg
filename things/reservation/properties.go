package reservation

import "github.com/dpb587/go-schemaorg"

var (
	// The service provider, service operator, or service performer; the goods
	// producer. Another party (a seller) may offer those services or goods on
	// behalf of the provider. A provider may also serve as the seller.
	Provider = schemaorg.NewProperty("provider")

	// Any membership in a frequent flyer, hotel loyalty program, etc. being applied
	// to the reservation.
	ProgramMembershipUsed = schemaorg.NewProperty("programMembershipUsed")

	// The current status of the reservation.
	ReservationStatus = schemaorg.NewProperty("reservationStatus")

	// The person or organization the reservation or ticket is for.
	UnderName = schemaorg.NewProperty("underName")

	// 'bookingAgent' is an out-dated term indicating a 'broker' that serves as a
	// booking agent.
	BookingAgent = schemaorg.NewProperty("bookingAgent")

	// The total price for the reservation or ticket, including applicable taxes,
	// shipping, etc.
	TotalPrice = schemaorg.NewProperty("totalPrice")

	// The thing -- flight, event, restaurant,etc. being reserved.
	ReservationFor = schemaorg.NewProperty("reservationFor")

	// The currency (in 3-letter ISO 4217 format) of the price or a price component,
	// when attached to <a class="localLink"
	// href="http://schema.org/PriceSpecification">PriceSpecification</a> and its
	// subtypes.
	PriceCurrency = schemaorg.NewProperty("priceCurrency")

	// An entity that arranges for an exchange between a buyer and a seller.  In
	// most cases a broker never acquires or releases ownership of a product or
	// service involved in an exchange.  If it is not clear whether an entity is a
	// broker, seller, or buyer, the latter two terms are preferred.
	Broker = schemaorg.NewProperty("broker")

	// The date and time the reservation was modified.
	ModifiedTime = schemaorg.NewProperty("modifiedTime")

	// The date and time the reservation was booked.
	BookingTime = schemaorg.NewProperty("bookingTime")

	// A unique identifier for the reservation.
	ReservationId = schemaorg.NewProperty("reservationId")

	// A ticket associated with the reservation.
	ReservedTicket = schemaorg.NewProperty("reservedTicket")
)
