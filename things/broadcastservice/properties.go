package broadcastservice

import "github.com/dpb587/go-schemaorg"

var (
	// The timezone in <a href="http://en.wikipedia.org/wiki/ISO_8601">ISO 8601
	// format</a> for which the service bases its broadcasts
	BroadcastTimezone = schemaorg.NewProperty("broadcastTimezone")

	// The type of screening or video broadcast used (e.g. IMAX, 3D, SD, HD, etc.).
	VideoFormat = schemaorg.NewProperty("videoFormat")

	// The media network(s) whose content is broadcast on this station.
	BroadcastAffiliateOf = schemaorg.NewProperty("broadcastAffiliateOf")

	// A broadcast service to which the broadcast service may belong to such as
	// regional variations of a national channel.
	ParentService = schemaorg.NewProperty("parentService")

	// The name displayed in the channel guide. For many US affiliates, it is the
	// network name.
	BroadcastDisplayName = schemaorg.NewProperty("broadcastDisplayName")

	// The area within which users can expect to reach the broadcast service.
	Area = schemaorg.NewProperty("area")

	// The organization owning or operating the broadcast service.
	Broadcaster = schemaorg.NewProperty("broadcaster")
)
