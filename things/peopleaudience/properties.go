package peopleaudience

import "github.com/dpb587/go-schemaorg"

var (
	// Maximal age recommended for viewing content.
	SuggestedMaxAge = schemaorg.NewProperty("suggestedMaxAge")

	// Audiences defined by a person's minimum age.
	RequiredMinAge = schemaorg.NewProperty("requiredMinAge")

	// Audiences defined by a person's gender.
	RequiredGender = schemaorg.NewProperty("requiredGender")

	// The gender of the person or audience.
	SuggestedGender = schemaorg.NewProperty("suggestedGender")

	// Audiences defined by a person's maximum age.
	RequiredMaxAge = schemaorg.NewProperty("requiredMaxAge")

	// Minimal age recommended for viewing content.
	SuggestedMinAge = schemaorg.NewProperty("suggestedMinAge")
)
