package contactpoint

import "github.com/dpb587/go-schemaorg"

var (
	// The geographic area where the service is provided.
	ServiceArea = schemaorg.NewProperty("serviceArea")

	// The geographic area where a service or offered item is provided.
	AreaServed = schemaorg.NewProperty("areaServed")

	// The hours during which this service or contact is available.
	HoursAvailable = schemaorg.NewProperty("hoursAvailable")

	// An option available on this contact point (e.g. a toll-free number or support
	// for hearing-impaired callers).
	ContactOption = schemaorg.NewProperty("contactOption")

	// A language someone may use with or at the item, service or place. Please use
	// one of the language codes from the <a
	// href="http://tools.ietf.org/html/bcp47">IETF BCP 47 standard</a>. See also <a
	// class="localLink" href="http://schema.org/inLanguage">inLanguage</a>
	AvailableLanguage = schemaorg.NewProperty("availableLanguage")

	// The telephone number.
	Telephone = schemaorg.NewProperty("telephone")

	// Email address.
	Email = schemaorg.NewProperty("email")

	// A person or organization can have different contact points, for different
	// purposes. For example, a sales contact point, a PR contact point and so on.
	// This property is used to specify the kind of contact point.
	ContactType = schemaorg.NewProperty("contactType")

	// The product or service this support contact point is related to (such as
	// product support for a particular product line). This can be a specific
	// product or product line (e.g. "iPhone") or a general category of products or
	// services (e.g. "smartphones").
	ProductSupported = schemaorg.NewProperty("productSupported")

	// The fax number.
	FaxNumber = schemaorg.NewProperty("faxNumber")
)
