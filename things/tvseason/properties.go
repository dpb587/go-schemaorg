package tvseason

import "github.com/dpb587/go-schemaorg"

var (
	// The TV series to which this episode or season belongs.
	PartOfTVSeries = schemaorg.NewProperty("partOfTVSeries")

	// The country of the principal offices of the production company or individual
	// responsible for the movie or program.
	CountryOfOrigin = schemaorg.NewProperty("countryOfOrigin")
)
