package techarticle

import "github.com/dpb587/go-schemaorg"

var (
	// Prerequisites needed to fulfill steps in article.
	Dependencies = schemaorg.NewProperty("dependencies")

	// Proficiency needed for this content; expected values: 'Beginner', 'Expert'.
	ProficiencyLevel = schemaorg.NewProperty("proficiencyLevel")
)
