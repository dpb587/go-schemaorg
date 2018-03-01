package offer

import "github.com/dpb587/go-schemaorg"

// // An offer to transfer some rights to an item or to provide a service — for
// example, an offer to sell tickets to an event, to rent the DVD of a movie, to
// stream a TV show over the internet, to repair a motorcycle, or to loan a
// book.</p>
// 
// <p>For <a
// href="http://www.gs1.org/barcodes/technical/idkeys/gtin">GTIN</a>-related
// fields, see <a
// href="http://www.gs1.org/barcodes/support/check_digit_calculator">Check Digit
// calculator</a> and <a
// href="http://www.gs1us.org/resources/standards/gtin-validation-guide">validation
// guide</a> from <a href="http://www.gs1.org/">GS1</a>.
var Type = schemaorg.NewDataType("http://schema.org", "Offer")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
