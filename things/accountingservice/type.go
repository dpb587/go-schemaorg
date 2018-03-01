package accountingservice

import "github.com/dpb587/go-schemaorg"

// // Accountancy business.</p>
// 
// <p>As a <a class="localLink"
// href="http://schema.org/LocalBusiness">LocalBusiness</a> it can be described
// as a <a class="localLink" href="http://schema.org/provider">provider</a> of
// one or more <a class="localLink"
// href="http://schema.org/Service">Service</a>(s).
var Type = schemaorg.NewDataType("http://schema.org", "AccountingService")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
