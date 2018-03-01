package howto

import "github.com/dpb587/go-schemaorg"

var (
	// The length of time it takes to prepare the items to be used in instructions
	// or a direction, in <a href="http://en.wikipedia.org/wiki/ISO_8601">ISO 8601
	// duration format</a>.
	PrepTime = schemaorg.NewProperty("prepTime")

	// A sub property of instrument. An object used (but not consumed) when
	// performing instructions or a direction.
	Tool = schemaorg.NewProperty("tool")

	// The steps in the form of a single item (text, document, video, etc.) or an
	// ordered list with HowToStep and/or HowToSection items.
	Steps = schemaorg.NewProperty("steps")

	// The length of time it takes to perform instructions or a direction (not
	// including time to prepare the supplies), in <a
	// href="http://en.wikipedia.org/wiki/ISO_8601">ISO 8601 duration format</a>.
	PerformTime = schemaorg.NewProperty("performTime")

	// The estimated cost of the supply or supplies consumed when performing
	// instructions.
	EstimatedCost = schemaorg.NewProperty("estimatedCost")

	// The total time required to perform instructions or a direction (including
	// time to prepare the supplies), in <a
	// href="http://en.wikipedia.org/wiki/ISO_8601">ISO 8601 duration format</a>.
	TotalTime = schemaorg.NewProperty("totalTime")

	// A sub-property of instrument. A supply consumed when performing instructions
	// or a direction.
	Supply = schemaorg.NewProperty("supply")

	// The quantity that results by performing instructions. For example, a paper
	// airplane, 10 personalized candles.
	Yield = schemaorg.NewProperty("yield")
)
