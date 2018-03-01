package alignmentobject

import "github.com/dpb587/go-schemaorg"

var (
	// The description of a node in an established educational framework.
	TargetDescription = schemaorg.NewProperty("targetDescription")

	// A category of alignment between the learning resource and the framework node.
	// Recommended values include: 'assesses', 'teaches', 'requires',
	// 'textComplexity', 'readingLevel', 'educationalSubject', and
	// 'educationalLevel'.
	AlignmentType = schemaorg.NewProperty("alignmentType")

	// The URL of a node in an established educational framework.
	TargetUrl = schemaorg.NewProperty("targetUrl")

	// The name of a node in an established educational framework.
	TargetName = schemaorg.NewProperty("targetName")

	// The framework to which the resource being described is aligned.
	EducationalFramework = schemaorg.NewProperty("educationalFramework")
)
