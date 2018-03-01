package event

import "github.com/dpb587/go-schemaorg"

var (
	// The subject matter of the content.
	About = schemaorg.NewProperty("about")

	// A person or organization that supports (sponsors) something through some kind
	// of financial contribution.
	Funder = schemaorg.NewProperty("funder")

	// A work featured in some event, e.g. exhibited in an ExhibitionEvent.
	//        Specific subproperties are available for workPerformed (e.g. a play),
	// or a workPresented (a Movie at a ScreeningEvent).
	WorkFeatured = schemaorg.NewProperty("workFeatured")

	// An intended audience, i.e. a group for whom something was created.
	Audience = schemaorg.NewProperty("audience")

	// The number of attendee places for an event that remain unallocated.
	RemainingAttendeeCapacity = schemaorg.NewProperty("remainingAttendeeCapacity")

	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors
	// can be associated with individual items or with a series, episode, clip.
	Actor = schemaorg.NewProperty("actor")

	// The main performer or performers of the event&#x2014;for example, a
	// presenter, musician, or actor.
	Performers = schemaorg.NewProperty("performers")

	// The end date and time of the item (in <a
	// href="http://en.wikipedia.org/wiki/ISO_8601">ISO 8601 date format</a>).
	EndDate = schemaorg.NewProperty("endDate")

	// The time admission will commence.
	DoorTime = schemaorg.NewProperty("doorTime")

	// A secondary contributor to the CreativeWork or Event.
	Contributor = schemaorg.NewProperty("contributor")

	// The total number of individuals that may attend an event or venue.
	MaximumAttendeeCapacity = schemaorg.NewProperty("maximumAttendeeCapacity")

	// The typical expected age range, e.g. '7-9', '11-'.
	TypicalAgeRange = schemaorg.NewProperty("typicalAgeRange")

	// An organizer of an Event.
	Organizer = schemaorg.NewProperty("organizer")

	// A person attending the event.
	Attendees = schemaorg.NewProperty("attendees")

	// The overall rating, based on a collection of reviews or ratings, of the item.
	AggregateRating = schemaorg.NewProperty("aggregateRating")

	// An Event that is part of this event. For example, a conference event includes
	// many presentations, each of which is a subEvent of the conference.
	SubEvent = schemaorg.NewProperty("subEvent")

	// Events that are a part of this event. For example, a conference event
	// includes many presentations, each subEvents of the conference.
	SubEvents = schemaorg.NewProperty("subEvents")

	// An offer to provide this item&#x2014;for example, an offer to sell a product,
	// rent the DVD of a movie, perform a service, or give away tickets to an event.
	Offers = schemaorg.NewProperty("offers")

	// The language of the content or performance or used in an action. Please use
	// one of the language codes from the <a
	// href="http://tools.ietf.org/html/bcp47">IETF BCP 47 standard</a>. See also <a
	// class="localLink"
	// href="http://schema.org/availableLanguage">availableLanguage</a>.
	InLanguage = schemaorg.NewProperty("inLanguage")

	// A person or organization attending the event.
	Attendee = schemaorg.NewProperty("attendee")

	// A work performed in some event, for example a play performed in a
	// TheaterEvent.
	WorkPerformed = schemaorg.NewProperty("workPerformed")

	// An eventStatus of an event represents its status; particularly useful when an
	// event is cancelled or rescheduled.
	EventStatus = schemaorg.NewProperty("eventStatus")

	// The start date and time of the item (in <a
	// href="http://en.wikipedia.org/wiki/ISO_8601">ISO 8601 date format</a>).
	StartDate = schemaorg.NewProperty("startDate")

	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an
	// event. Directors can be associated with individual items or with a series,
	// episode, clip.
	Director = schemaorg.NewProperty("director")

	// An event that this event is a part of. For example, a collection of
	// individual music performances might each have a music festival as their
	// superEvent.
	SuperEvent = schemaorg.NewProperty("superEvent")

	// The duration of the item (movie, audio recording, event, etc.) in <a
	// href="http://en.wikipedia.org/wiki/ISO_8601">ISO 8601 date format</a>.
	Duration = schemaorg.NewProperty("duration")

	// Organization or person who adapts a creative work to different languages,
	// regional differences and technical requirements of a target market, or that
	// translates during some event.
	Translator = schemaorg.NewProperty("translator")

	// Used in conjunction with eventStatus for rescheduled or cancelled events.
	// This property contains the previously scheduled start date. For rescheduled
	// events, the startDate property should be used for the newly scheduled start
	// date. In the (rare) case of an event that has been postponed and rescheduled
	// multiple times, this field may be repeated.
	PreviousStartDate = schemaorg.NewProperty("previousStartDate")

	// A review of the item.
	Review = schemaorg.NewProperty("review")

	// A person or organization that supports a thing through a pledge, promise, or
	// financial contribution. e.g. a sponsor of a Medical Study or a corporate
	// sponsor of an event.
	Sponsor = schemaorg.NewProperty("sponsor")

	// The location of for example where the event is happening, an organization is
	// located, or where an action takes place.
	Location = schemaorg.NewProperty("location")

	// The CreativeWork that captured all or part of this Event.
	RecordedIn = schemaorg.NewProperty("recordedIn")

	// The person or organization who wrote a composition, or who is the composer of
	// a work performed at some event.
	Composer = schemaorg.NewProperty("composer")

	// A flag to signal that the item, event, or place is accessible for free.
	IsAccessibleForFree = schemaorg.NewProperty("isAccessibleForFree")

	// A performer at the event&#x2014;for example, a presenter, musician, musical
	// group or actor.
	Performer = schemaorg.NewProperty("performer")
)
