package lodgingbusiness

import "github.com/dpb587/go-schemaorg"

var (
	// An intended audience, i.e. a group for whom something was created.
	Audience = schemaorg.NewProperty("audience")

	// The earliest someone may check into a lodging establishment.
	CheckinTime = schemaorg.NewProperty("checkinTime")

	// Indicates whether pets are allowed to enter the accommodation or lodging
	// business. More detailed information can be put in a text value.
	PetsAllowed = schemaorg.NewProperty("petsAllowed")

	// A language someone may use with or at the item, service or place. Please use
	// one of the language codes from the <a
	// href="http://tools.ietf.org/html/bcp47">IETF BCP 47 standard</a>. See also <a
	// class="localLink" href="http://schema.org/inLanguage">inLanguage</a>
	AvailableLanguage = schemaorg.NewProperty("availableLanguage")

	// An amenity feature (e.g. a characteristic or service) of the Accommodation.
	// This generic property does not make a statement about whether the feature is
	// included in an offer for the main accommodation or available at extra costs.
	AmenityFeature = schemaorg.NewProperty("amenityFeature")

	// An official rating for a lodging business or food establishment, e.g. from
	// national associations or standards bodies. Use the author property to
	// indicate the rating organization, e.g. as an Organization with name such as
	// (e.g. HOTREC, DEHOGA, WHR, or Hotelstars).
	StarRating = schemaorg.NewProperty("starRating")

	// The latest someone may check out of a lodging establishment.
	CheckoutTime = schemaorg.NewProperty("checkoutTime")
)
