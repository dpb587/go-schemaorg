package place

import "github.com/dpb587/go-schemaorg"

var (
	// A photograph of this place.
	Photo = schemaorg.NewProperty("photo")

	// The opening hours of a certain place.
	OpeningHoursSpecification = schemaorg.NewProperty("openingHoursSpecification")

	// Upcoming or past events associated with this place or organization.
	Events = schemaorg.NewProperty("events")

	// Indicates whether it is allowed to smoke in the place, e.g. in the
	// restaurant, hotel or hotel room.
	SmokingAllowed = schemaorg.NewProperty("smokingAllowed")

	// The <a href="http://www.gs1.org/gln">Global Location Number</a> (GLN,
	// sometimes also referred to as International Location Number or ILN) of the
	// respective organization, person, or place. The GLN is a 13-digit number used
	// to identify parties and physical locations.
	GlobalLocationNumber = schemaorg.NewProperty("globalLocationNumber")

	// The total number of individuals that may attend an event or venue.
	MaximumAttendeeCapacity = schemaorg.NewProperty("maximumAttendeeCapacity")

	// Review of the item.
	Reviews = schemaorg.NewProperty("reviews")

	// The overall rating, based on a collection of reviews or ratings, of the item.
	AggregateRating = schemaorg.NewProperty("aggregateRating")

	// Photographs of this place.
	Photos = schemaorg.NewProperty("photos")

	// A URL to a map of the place.
	Map = schemaorg.NewProperty("map")

	// A short textual code (also called "store code") that uniquely identifies a
	// place of business. The code is typically assigned by the parentOrganization
	// and used in structured URLs.</p>
	// 
	// <p>For example, in the URL
	// http://www.starbucks.co.uk/store-locator/etc/detail/3047 the code "3047" is a
	// branchCode for a particular branch.
	BranchCode = schemaorg.NewProperty("branchCode")

	// A URL to a map of the place.
	HasMap = schemaorg.NewProperty("hasMap")

	// A property-value pair representing an additional characteristics of the
	// entitity, e.g. a product feature or another characteristic for which there is
	// no matching property in schema.org.</p>
	// 
	// <p>Note: Publishers should be aware that applications designed to use
	// specific schema.org properties (e.g. http://schema.org/width,
	// http://schema.org/color, http://schema.org/gtin13, ...) will typically expect
	// such data to be provided using those properties, rather than using the
	// generic property/value mechanism.
	AdditionalProperty = schemaorg.NewProperty("additionalProperty")

	// Physical address of the item.
	Address = schemaorg.NewProperty("address")

	// The special opening hours of a certain place.</p>
	// 
	// <p>Use this to explicitly override general opening hours brought in scope by
	// <a class="localLink"
	// href="http://schema.org/openingHoursSpecification">openingHoursSpecification</a>
	// or <a class="localLink"
	// href="http://schema.org/openingHours">openingHours</a>.
	SpecialOpeningHoursSpecification = schemaorg.NewProperty("specialOpeningHoursSpecification")

	// An amenity feature (e.g. a characteristic or service) of the Accommodation.
	// This generic property does not make a statement about whether the feature is
	// included in an offer for the main accommodation or available at extra costs.
	AmenityFeature = schemaorg.NewProperty("amenityFeature")

	// An associated logo.
	Logo = schemaorg.NewProperty("logo")

	// The telephone number.
	Telephone = schemaorg.NewProperty("telephone")

	// The geo coordinates of the place.
	Geo = schemaorg.NewProperty("geo")

	// The basic containment relation between a place and one that contains it.
	ContainedInPlace = schemaorg.NewProperty("containedInPlace")

	// A review of the item.
	Review = schemaorg.NewProperty("review")

	// A flag to signal that the <a class="localLink"
	// href="http://schema.org/Place">Place</a> is open to public visitors.  If this
	// property is omitted there is no assumed default boolean value
	PublicAccess = schemaorg.NewProperty("publicAccess")

	// Upcoming or past event associated with this place, organization, or action.
	Event = schemaorg.NewProperty("event")

	// The basic containment relation between a place and another that it contains.
	ContainsPlace = schemaorg.NewProperty("containsPlace")

	// The International Standard of Industrial Classification of All Economic
	// Activities (ISIC), Revision 4 code for a particular organization, business
	// person, or place.
	IsicV4 = schemaorg.NewProperty("isicV4")

	// A URL to a map of the place.
	Maps = schemaorg.NewProperty("maps")

	// The fax number.
	FaxNumber = schemaorg.NewProperty("faxNumber")

	// A flag to signal that the item, event, or place is accessible for free.
	IsAccessibleForFree = schemaorg.NewProperty("isAccessibleForFree")

	// The basic containment relation between a place and one that contains it.
	ContainedIn = schemaorg.NewProperty("containedIn")
)
