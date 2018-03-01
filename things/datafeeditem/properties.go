package datafeeditem

import "github.com/dpb587/go-schemaorg"

var (
	// The date on which the CreativeWork was most recently modified or when the
	// item's entry was modified within a DataFeed.
	DateModified = schemaorg.NewProperty("dateModified")

	// The datetime the item was removed from the DataFeed.
	DateDeleted = schemaorg.NewProperty("dateDeleted")

	// The date on which the CreativeWork was created or the item was added to a
	// DataFeed.
	DateCreated = schemaorg.NewProperty("dateCreated")

	// An entity represented by an entry in a list or data feed (e.g. an 'artist' in
	// a list of 'artists')’.
	Item = schemaorg.NewProperty("item")
)
