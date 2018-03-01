package foodestablishment

import "github.com/dpb587/go-schemaorg"

var (
	// The cuisine of the restaurant.
	ServesCuisine = schemaorg.NewProperty("servesCuisine")

	// Either the actual menu as a structured representation, as text, or a URL of
	// the menu.
	HasMenu = schemaorg.NewProperty("hasMenu")

	// Either the actual menu as a structured representation, as text, or a URL of
	// the menu.
	Menu = schemaorg.NewProperty("menu")

	// An official rating for a lodging business or food establishment, e.g. from
	// national associations or standards bodies. Use the author property to
	// indicate the rating organization, e.g. as an Organization with name such as
	// (e.g. HOTREC, DEHOGA, WHR, or Hotelstars).
	StarRating = schemaorg.NewProperty("starRating")

	// Indicates whether a FoodEstablishment accepts reservations. Values can be
	// Boolean, an URL at which reservations can be made or (for backwards
	// compatibility) the strings <code>Yes</code> or <code>No</code>.
	AcceptsReservations = schemaorg.NewProperty("acceptsReservations")
)
