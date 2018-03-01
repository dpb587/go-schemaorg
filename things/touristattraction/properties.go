package touristattraction

import "github.com/dpb587/go-schemaorg"

var (
	// Attraction suitable for type(s) of tourist. eg. Children, visitors from a
	// particular country, etc.
	TouristType = schemaorg.NewProperty("touristType")

	// A language someone may use with or at the item, service or place. Please use
	// one of the language codes from the <a
	// href="http://tools.ietf.org/html/bcp47">IETF BCP 47 standard</a>. See also <a
	// class="localLink" href="http://schema.org/inLanguage">inLanguage</a>
	AvailableLanguage = schemaorg.NewProperty("availableLanguage")
)
