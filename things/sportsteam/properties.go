package sportsteam

import "github.com/dpb587/go-schemaorg"

var (
	// A person that acts as performing member of a sports team; a player as opposed
	// to a coach.
	Athlete = schemaorg.NewProperty("athlete")

	// A person that acts in a coaching role for a sports team.
	Coach = schemaorg.NewProperty("coach")
)
