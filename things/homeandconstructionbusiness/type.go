package homeandconstructionbusiness

import "github.com/dpb587/go-schemaorg"

// // A construction business.</p>
// 
// <p>A HomeAndConstructionBusiness is a <a class="localLink"
// href="http://schema.org/LocalBusiness">LocalBusiness</a> that provides
// services around homes and buildings.</p>
// 
// <p>As a <a class="localLink"
// href="http://schema.org/LocalBusiness">LocalBusiness</a> it can be described
// as a <a class="localLink" href="http://schema.org/provider">provider</a> of
// one or more <a class="localLink"
// href="http://schema.org/Service">Service</a>(s).
var Type = schemaorg.NewDataType("http://schema.org", "HomeAndConstructionBusiness")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
