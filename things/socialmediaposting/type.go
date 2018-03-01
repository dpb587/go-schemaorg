package socialmediaposting

import "github.com/dpb587/go-schemaorg"

// // A post to a social media platform, including blog posts, tweets, Facebook
// posts, etc.
var Type = schemaorg.NewDataType("http://schema.org", "SocialMediaPosting")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
