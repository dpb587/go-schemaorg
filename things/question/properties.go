package question

import "github.com/dpb587/go-schemaorg"

var (
	// The number of upvotes this question, answer or comment has received from the
	// community.
	UpvoteCount = schemaorg.NewProperty("upvoteCount")

	// The number of answers this question has received.
	AnswerCount = schemaorg.NewProperty("answerCount")

	// The answer that has been accepted as best, typically on a Question/Answer
	// site. Sites vary in their selection mechanisms, e.g. drawing on community
	// opinion and/or the view of the Question author.
	AcceptedAnswer = schemaorg.NewProperty("acceptedAnswer")

	// An answer (possibly one of several, possibly incorrect) to a Question, e.g.
	// on a Question/Answer site.
	SuggestedAnswer = schemaorg.NewProperty("suggestedAnswer")

	// The number of downvotes this question, answer or comment has received from
	// the community.
	DownvoteCount = schemaorg.NewProperty("downvoteCount")
)
