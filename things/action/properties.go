package action

import "github.com/dpb587/go-schemaorg"

var (
	// The result produced in the action. e.g. John wrote <em>a book</em>.
	Result = schemaorg.NewProperty("result")

	// Indicates the current disposition of the Action.
	ActionStatus = schemaorg.NewProperty("actionStatus")

	// Indicates a target EntryPoint for an Action.
	Target = schemaorg.NewProperty("target")

	// The direct performer or driver of the action (animate or inanimate). e.g.
	// <em>John</em> wrote a book.
	Agent = schemaorg.NewProperty("agent")

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

	// The object that helped the agent perform the action. e.g. John wrote a book
	// with <em>a pen</em>.
	Instrument = schemaorg.NewProperty("instrument")

	// Other co-agents that participated in the action indirectly. e.g. John wrote a
	// book with <em>Steve</em>.
	Participant = schemaorg.NewProperty("participant")

	// The object upon which the action is carried out, whose state is kept intact
	// or changed. Also known as the semantic roles patient, affected or undergoer
	// (which change their state) or theme (which doesn't). e.g. John read <em>a
	// book</em>.
	Object = schemaorg.NewProperty("object")

	// For failed actions, more information on the cause of the failure.
	Error = schemaorg.NewProperty("error")

	// The location of for example where the event is happening, an organization is
	// located, or where an action takes place.
	Location = schemaorg.NewProperty("location")
)
