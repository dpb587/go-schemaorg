package discussionforumposting

import "github.com/dpb587/go-schemaorg"

// // A posting to a discussion forum.
var Type = schemaorg.NewDataType("http://schema.org", "DiscussionForumPosting")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
