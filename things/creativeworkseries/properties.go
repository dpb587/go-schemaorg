package creativeworkseries

import "github.com/dpb587/go-schemaorg"

var (
	// The end date and time of the item (in <a
	// href="http://en.wikipedia.org/wiki/ISO_8601">ISO 8601 date format</a>).
	EndDate = schemaorg.NewProperty("endDate")

	// The International Standard Serial Number (ISSN) that identifies this serial
	// publication. You can repeat this property to identify different formats of,
	// or the linking ISSN (ISSN-L) for, this serial publication.
	Issn = schemaorg.NewProperty("issn")

	// The start date and time of the item (in <a
	// href="http://en.wikipedia.org/wiki/ISO_8601">ISO 8601 date format</a>).
	StartDate = schemaorg.NewProperty("startDate")
)
