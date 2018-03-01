package geoshape

import "github.com/dpb587/go-schemaorg"

var (
	// A line is a point-to-point path consisting of two or more points. A line is
	// expressed as a series of two or more point objects separated by space.
	Line = schemaorg.NewProperty("line")

	// Physical address of the item.
	Address = schemaorg.NewProperty("address")

	// A circle is the circular region of a specified radius centered at a specified
	// latitude and longitude. A circle is expressed as a pair followed by a radius
	// in meters.
	Circle = schemaorg.NewProperty("circle")

	// A box is the area enclosed by the rectangle formed by two points. The first
	// point is the lower corner, the second point is the upper corner. A box is
	// expressed as two points separated by a space character.
	Box = schemaorg.NewProperty("box")

	// The country. For example, USA. You can also provide the two-letter <a
	// href="http://en.wikipedia.org/wiki/ISO_3166-1">ISO 3166-1 alpha-2 country
	// code</a>.
	AddressCountry = schemaorg.NewProperty("addressCountry")

	// The postal code. For example, 94043.
	PostalCode = schemaorg.NewProperty("postalCode")

	// The elevation of a location (<a
	// href="https://en.wikipedia.org/wiki/World_Geodetic_System">WGS 84</a>).
	Elevation = schemaorg.NewProperty("elevation")

	// A polygon is the area enclosed by a point-to-point path for which the
	// starting and ending points are the same. A polygon is expressed as a series
	// of four or more space delimited points where the first and final points are
	// identical.
	Polygon = schemaorg.NewProperty("polygon")
)
