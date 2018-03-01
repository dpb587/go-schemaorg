package audioobject

import "github.com/dpb587/go-schemaorg"

var (
	// If this MediaObject is an AudioObject or VideoObject, the transcript of that
	// object.
	Transcript = schemaorg.NewProperty("transcript")
)
