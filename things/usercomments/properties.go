package usercomments

import "github.com/dpb587/go-schemaorg"

var (
	// The text of the UserComment.
	CommentText = schemaorg.NewProperty("commentText")

	// Specifies the CreativeWork associated with the UserComment.
	Discusses = schemaorg.NewProperty("discusses")

	// The time at which the UserComment was made.
	CommentTime = schemaorg.NewProperty("commentTime")

	// The creator/author of this CreativeWork. This is the same as the Author
	// property for CreativeWork.
	Creator = schemaorg.NewProperty("creator")

	// The URL at which a reply may be posted to the specified UserComment.
	ReplyToUrl = schemaorg.NewProperty("replyToUrl")
)
