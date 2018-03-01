package geocoordinates

import "github.com/dpb587/go-schemaorg"

var (
	// The latitude of a location. For example <code>37.42242</code> (<a
	// href="https://en.wikipedia.org/wiki/World_Geodetic_System">WGS 84</a>).
	Latitude = schemaorg.NewProperty("latitude")

	// The longitude of a location. For example <code>-122.08585</code> (<a
	// href="https://en.wikipedia.org/wiki/World_Geodetic_System">WGS 84</a>).
	Longitude = schemaorg.NewProperty("longitude")

	// Physical address of the item.
	Address = schemaorg.NewProperty("address")

	// The country. For example, USA. You can also provide the two-letter <a
	// href="http://en.wikipedia.org/wiki/ISO_3166-1">ISO 3166-1 alpha-2 country
	// code</a>.
	AddressCountry = schemaorg.NewProperty("addressCountry")

	// The postal code. For example, 94043.
	PostalCode = schemaorg.NewProperty("postalCode")

	// The elevation of a location (<a
	// href="https://en.wikipedia.org/wiki/World_Geodetic_System">WGS 84</a>).
	Elevation = schemaorg.NewProperty("elevation")
)
