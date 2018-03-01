package interactioncounter

import "github.com/dpb587/go-schemaorg"

var (
	// The WebSite or SoftwareApplication where the interactions took place.
	InteractionService = schemaorg.NewProperty("interactionService")

	// The number of interactions for the CreativeWork using the WebSite or
	// SoftwareApplication.
	UserInteractionCount = schemaorg.NewProperty("userInteractionCount")

	// The Action representing the type of interaction. For up votes, +1s, etc. use
	// <a class="localLink" href="http://schema.org/LikeAction">LikeAction</a>. For
	// down votes use <a class="localLink"
	// href="http://schema.org/DislikeAction">DislikeAction</a>. Otherwise, use the
	// most specific Action.
	InteractionType = schemaorg.NewProperty("interactionType")
)
