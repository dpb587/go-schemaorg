package bookmarkaction

import "github.com/dpb587/go-schemaorg"

// // An agent bookmarks/flags/labels/tags/marks an object.
var Type = schemaorg.NewDataType("http://schema.org", "BookmarkAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
