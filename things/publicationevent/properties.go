package publicationevent

import "github.com/dpb587/go-schemaorg"

var (
	// A broadcast service associated with the publication event.
	PublishedOn = schemaorg.NewProperty("publishedOn")

	// A flag to signal that the item, event, or place is accessible for free.
	Free = schemaorg.NewProperty("free")

	// A flag to signal that the item, event, or place is accessible for free.
	IsAccessibleForFree = schemaorg.NewProperty("isAccessibleForFree")
)
