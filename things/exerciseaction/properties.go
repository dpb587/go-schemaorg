package exerciseaction

import "github.com/dpb587/go-schemaorg"

var (
	// A sub property of location. The course where this action was taken.
	Course = schemaorg.NewProperty("course")

	// A sub property of participant. The sports team that participated on this
	// action.
	SportsTeam = schemaorg.NewProperty("sportsTeam")

	// A sub property of location. The sports event where this action occurred.
	SportsEvent = schemaorg.NewProperty("sportsEvent")

	// The distance travelled, e.g. exercising or travelling.
	Distance = schemaorg.NewProperty("distance")

	// A sub property of participant. The opponent on this action.
	Opponent = schemaorg.NewProperty("opponent")

	// A sub property of location. The sports activity location where this action
	// occurred.
	SportsActivityLocation = schemaorg.NewProperty("sportsActivityLocation")

	// A sub property of location. The final location of the object or the agent
	// after the action.
	ToLocation = schemaorg.NewProperty("toLocation")

	// A sub property of location. The original location of the object or the agent
	// before the action.
	FromLocation = schemaorg.NewProperty("fromLocation")

	// A sub property of location. The course where this action was taken.
	ExerciseCourse = schemaorg.NewProperty("exerciseCourse")
)
