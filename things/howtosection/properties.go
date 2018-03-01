package howtosection

import "github.com/dpb587/go-schemaorg"

var (
	// The steps in the form of a single item (text, document, video, etc.) or an
	// ordered list with HowToStep and/or HowToSection items.
	Steps = schemaorg.NewProperty("steps")
)
