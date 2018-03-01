package professionalservice

import "github.com/dpb587/go-schemaorg"

// // Original definition: "provider of professional services."</p>
// 
// <p>The general <a class="localLink"
// href="http://schema.org/ProfessionalService">ProfessionalService</a> type for
// local businesses was deprecated due to confusion with <a class="localLink"
// href="http://schema.org/Service">Service</a>. For reference, the types that
// it included were: <a class="localLink"
// href="http://schema.org/Dentist">Dentist</a>,
//         <a class="localLink"
// href="http://schema.org/AccountingService">AccountingService</a>, <a
// class="localLink" href="http://schema.org/Attorney">Attorney</a>, <a
// class="localLink" href="http://schema.org/Notary">Notary</a>, as well as
// types for several kinds of <a class="localLink"
// href="http://schema.org/HomeAndConstructionBusiness">HomeAndConstructionBusiness</a>:
// <a class="localLink" href="http://schema.org/Electrician">Electrician</a>, <a
// class="localLink"
// href="http://schema.org/GeneralContractor">GeneralContractor</a>,
//         <a class="localLink"
// href="http://schema.org/HousePainter">HousePainter</a>, <a class="localLink"
// href="http://schema.org/Locksmith">Locksmith</a>, <a class="localLink"
// href="http://schema.org/Plumber">Plumber</a>, <a class="localLink"
// href="http://schema.org/RoofingContractor">RoofingContractor</a>. <a
// class="localLink" href="http://schema.org/LegalService">LegalService</a> was
// introduced as a more inclusive supertype of <a class="localLink"
// href="http://schema.org/Attorney">Attorney</a>.
var Type = schemaorg.NewDataType("http://schema.org", "ProfessionalService")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
