package screeningevent

import "github.com/dpb587/go-schemaorg"

var (
	// The type of screening or video broadcast used (e.g. IMAX, 3D, SD, HD, etc.).
	VideoFormat = schemaorg.NewProperty("videoFormat")

	// Languages in which subtitles/captions are available, in <a
	// href="http://tools.ietf.org/html/bcp47">IETF BCP 47 standard format</a>.
	SubtitleLanguage = schemaorg.NewProperty("subtitleLanguage")

	// The movie presented during this event.
	WorkPresented = schemaorg.NewProperty("workPresented")
)
