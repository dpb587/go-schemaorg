package sportsevent

import "github.com/dpb587/go-schemaorg"

var (
	// The away team in a sports event.
	AwayTeam = schemaorg.NewProperty("awayTeam")

	// A competitor in a sports event.
	Competitor = schemaorg.NewProperty("competitor")

	// The home team in a sports event.
	HomeTeam = schemaorg.NewProperty("homeTeam")
)
