package touristattraction

import "github.com/dpb587/go-schemaorg"

// // A tourist attraction.  In principle any Thing can be a <a class="localLink"
// href="http://schema.org/TouristAttraction">TouristAttraction</a>, from a <a
// class="localLink" href="http://schema.org/Mountain">Mountain</a> and <a
// class="localLink"
// href="http://schema.org/LandmarksOrHistoricalBuildings">LandmarksOrHistoricalBuildings</a>
// to a <a class="localLink"
// href="http://schema.org/LocalBusiness">LocalBusiness</a>.  This Type can be
// used on its own to describe a general <a class="localLink"
// href="http://schema.org/TourstAttraction">TourstAttraction</a>, or be used as
// an <a class="localLink"
// href="http://schema.org/additionalType">additionalType</a> to add tourist
// attraction properties to any other type.  (See examples below)
var Type = schemaorg.NewDataType("http://schema.org", "TouristAttraction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
