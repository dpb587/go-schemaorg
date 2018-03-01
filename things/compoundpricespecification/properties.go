package compoundpricespecification

import "github.com/dpb587/go-schemaorg"

var (
	// This property links to all <a class="localLink"
	// href="http://schema.org/UnitPriceSpecification">UnitPriceSpecification</a>
	// nodes that apply in parallel for the <a class="localLink"
	// href="http://schema.org/CompoundPriceSpecification">CompoundPriceSpecification</a>
	// node.
	PriceComponent = schemaorg.NewProperty("priceComponent")
)
