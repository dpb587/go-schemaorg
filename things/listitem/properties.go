package listitem

import "github.com/dpb587/go-schemaorg"

var (
	// A link to the ListItem that follows the current one.
	NextItem = schemaorg.NewProperty("nextItem")

	// The position of an item in a series or sequence of items.
	Position = schemaorg.NewProperty("position")

	// A link to the ListItem that preceeds the current one.
	PreviousItem = schemaorg.NewProperty("previousItem")

	// An entity represented by an entry in a list or data feed (e.g. an 'artist' in
	// a list of 'artists')’.
	Item = schemaorg.NewProperty("item")
)
