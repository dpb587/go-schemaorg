package broadcastchannel

import "github.com/dpb587/go-schemaorg"

var (
	// The unique address by which the BroadcastService can be identified in a
	// provider lineup. In US, this is typically a number.
	BroadcastChannelId = schemaorg.NewProperty("broadcastChannelId")

	// The type of service required to have access to the channel (e.g. Standard or
	// Premium).
	BroadcastServiceTier = schemaorg.NewProperty("broadcastServiceTier")

	// The CableOrSatelliteService offering the channel.
	InBroadcastLineup = schemaorg.NewProperty("inBroadcastLineup")

	// The BroadcastService offered on this channel.
	ProvidesBroadcastService = schemaorg.NewProperty("providesBroadcastService")

	// Genre of the creative work, broadcast channel or group.
	Genre = schemaorg.NewProperty("genre")
)
