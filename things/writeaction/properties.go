package writeaction

import "github.com/dpb587/go-schemaorg"

var (
	// The language of the content or performance or used in an action. Please use
	// one of the language codes from the <a
	// href="http://tools.ietf.org/html/bcp47">IETF BCP 47 standard</a>. See also <a
	// class="localLink"
	// href="http://schema.org/availableLanguage">availableLanguage</a>.
	InLanguage = schemaorg.NewProperty("inLanguage")

	// A sub property of instrument. The language used on this action.
	Language = schemaorg.NewProperty("language")
)
