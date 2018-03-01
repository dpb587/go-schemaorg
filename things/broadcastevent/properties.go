package broadcastevent

import "github.com/dpb587/go-schemaorg"

var (
	// The type of screening or video broadcast used (e.g. IMAX, 3D, SD, HD, etc.).
	VideoFormat = schemaorg.NewProperty("videoFormat")

	// True is the broadcast is of a live event.
	IsLiveBroadcast = schemaorg.NewProperty("isLiveBroadcast")

	// The event being broadcast such as a sporting event or awards ceremony.
	BroadcastOfEvent = schemaorg.NewProperty("broadcastOfEvent")
)
