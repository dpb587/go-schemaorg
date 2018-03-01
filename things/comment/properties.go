package comment

import "github.com/dpb587/go-schemaorg"

var (
	// The number of upvotes this question, answer or comment has received from the
	// community.
	UpvoteCount = schemaorg.NewProperty("upvoteCount")

	// The parent of a question, answer or item in general.
	ParentItem = schemaorg.NewProperty("parentItem")

	// The number of downvotes this question, answer or comment has received from
	// the community.
	DownvoteCount = schemaorg.NewProperty("downvoteCount")
)
