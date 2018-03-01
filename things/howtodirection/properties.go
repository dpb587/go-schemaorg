package howtodirection

import "github.com/dpb587/go-schemaorg"

var (
	// A media object representing the circumstances before performing this
	// direction.
	BeforeMedia = schemaorg.NewProperty("beforeMedia")

	// The length of time it takes to prepare the items to be used in instructions
	// or a direction, in <a href="http://en.wikipedia.org/wiki/ISO_8601">ISO 8601
	// duration format</a>.
	PrepTime = schemaorg.NewProperty("prepTime")

	// A sub property of instrument. An object used (but not consumed) when
	// performing instructions or a direction.
	Tool = schemaorg.NewProperty("tool")

	// The length of time it takes to perform instructions or a direction (not
	// including time to prepare the supplies), in <a
	// href="http://en.wikipedia.org/wiki/ISO_8601">ISO 8601 duration format</a>.
	PerformTime = schemaorg.NewProperty("performTime")

	// The total time required to perform instructions or a direction (including
	// time to prepare the supplies), in <a
	// href="http://en.wikipedia.org/wiki/ISO_8601">ISO 8601 duration format</a>.
	TotalTime = schemaorg.NewProperty("totalTime")

	// A sub-property of instrument. A supply consumed when performing instructions
	// or a direction.
	Supply = schemaorg.NewProperty("supply")

	// A media object representing the circumstances after performing this
	// direction.
	AfterMedia = schemaorg.NewProperty("afterMedia")

	// A media object representing the circumstances while performing this
	// direction.
	DuringMedia = schemaorg.NewProperty("duringMedia")
)
