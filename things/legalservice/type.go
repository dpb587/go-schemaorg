package legalservice

import "github.com/dpb587/go-schemaorg"

// // A LegalService is a business that provides legally-oriented services, advice
// and representation, e.g. law firms.</p>
// 
// <p>As a <a class="localLink"
// href="http://schema.org/LocalBusiness">LocalBusiness</a> it can be described
// as a <a class="localLink" href="http://schema.org/provider">provider</a> of
// one or more <a class="localLink"
// href="http://schema.org/Service">Service</a>(s).
var Type = schemaorg.NewDataType("http://schema.org", "LegalService")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
