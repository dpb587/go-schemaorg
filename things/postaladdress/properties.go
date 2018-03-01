package postaladdress

import "github.com/dpb587/go-schemaorg"

var (
	// The post office box number for PO box addresses.
	PostOfficeBoxNumber = schemaorg.NewProperty("postOfficeBoxNumber")

	// The street address. For example, 1600 Amphitheatre Pkwy.
	StreetAddress = schemaorg.NewProperty("streetAddress")

	// The country. For example, USA. You can also provide the two-letter <a
	// href="http://en.wikipedia.org/wiki/ISO_3166-1">ISO 3166-1 alpha-2 country
	// code</a>.
	AddressCountry = schemaorg.NewProperty("addressCountry")

	// The region. For example, CA.
	AddressRegion = schemaorg.NewProperty("addressRegion")

	// The postal code. For example, 94043.
	PostalCode = schemaorg.NewProperty("postalCode")

	// The locality. For example, Mountain View.
	AddressLocality = schemaorg.NewProperty("addressLocality")
)
