package foodestablishmentreservation

import "github.com/dpb587/go-schemaorg"

var (
	// The startTime of something. For a reserved event or service (e.g.
	// FoodEstablishmentReservation), the time that it is expected to start. For
	// actions that span a period of time, when the action was performed. e.g. John
	// wrote a book from <em>January</em> to December.</p>
	// 
	// <p>Note that Event uses startDate/endDate instead of startTime/endTime, even
	// when describing dates with times. This situation may be clarified in future
	// revisions.
	StartTime = schemaorg.NewProperty("startTime")

	// The endTime of something. For a reserved event or service (e.g.
	// FoodEstablishmentReservation), the time that it is expected to end. For
	// actions that span a period of time, when the action was performed. e.g. John
	// wrote a book from January to <em>December</em>.</p>
	// 
	// <p>Note that Event uses startDate/endDate instead of startTime/endTime, even
	// when describing dates with times. This situation may be clarified in future
	// revisions.
	EndTime = schemaorg.NewProperty("endTime")

	// Number of people the reservation should accommodate.
	PartySize = schemaorg.NewProperty("partySize")
)
