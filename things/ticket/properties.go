package ticket

import "github.com/dpb587/go-schemaorg"

var (
	// The person or organization the reservation or ticket is for.
	UnderName = schemaorg.NewProperty("underName")

	// The date the ticket was issued.
	DateIssued = schemaorg.NewProperty("dateIssued")

	// The total price for the reservation or ticket, including applicable taxes,
	// shipping, etc.
	TotalPrice = schemaorg.NewProperty("totalPrice")

	// The currency (in 3-letter ISO 4217 format) of the price or a price component,
	// when attached to <a class="localLink"
	// href="http://schema.org/PriceSpecification">PriceSpecification</a> and its
	// subtypes.
	PriceCurrency = schemaorg.NewProperty("priceCurrency")

	// The unique identifier for the ticket.
	TicketNumber = schemaorg.NewProperty("ticketNumber")

	// The organization issuing the ticket or permit.
	IssuedBy = schemaorg.NewProperty("issuedBy")

	// Reference to an asset (e.g., Barcode, QR code image or PDF) usable for
	// entrance.
	TicketToken = schemaorg.NewProperty("ticketToken")

	// The seat associated with the ticket.
	TicketedSeat = schemaorg.NewProperty("ticketedSeat")
)
