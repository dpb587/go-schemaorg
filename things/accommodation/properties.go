package accommodation

import "github.com/dpb587/go-schemaorg"

var (
	// The size of the accommodation, e.g. in square meter or squarefoot.
	// Typical unit code(s): MTK for square meter, FTK for square foot, or YDK for
	// square yard
	FloorSize = schemaorg.NewProperty("floorSize")

	// Indicates whether pets are allowed to enter the accommodation or lodging
	// business. More detailed information can be put in a text value.
	PetsAllowed = schemaorg.NewProperty("petsAllowed")

	// An amenity feature (e.g. a characteristic or service) of the Accommodation.
	// This generic property does not make a statement about whether the feature is
	// included in an offer for the main accommodation or available at extra costs.
	AmenityFeature = schemaorg.NewProperty("amenityFeature")

	// The number of rooms (excluding bathrooms and closets) of the acccommodation
	// or lodging business.
	// Typical unit code(s): ROM for room or C62 for no unit. The type of room can
	// be put in the unitText property of the QuantitativeValue.
	NumberOfRooms = schemaorg.NewProperty("numberOfRooms")

	// Indications regarding the permitted usage of the accommodation.
	PermittedUsage = schemaorg.NewProperty("permittedUsage")
)
