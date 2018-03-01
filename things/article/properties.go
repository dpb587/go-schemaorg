package article

import "github.com/dpb587/go-schemaorg"

var (
	// Any description of pages that is not separated into pageStart and pageEnd;
	// for example, "1-6, 9, 55" or "10-12, 46-49".
	Pagination = schemaorg.NewProperty("pagination")

	// The page on which the work ends; for example "138" or "xvi".
	PageEnd = schemaorg.NewProperty("pageEnd")

	// Articles may belong to one or more 'sections' in a magazine or newspaper,
	// such as Sports, Lifestyle, etc.
	ArticleSection = schemaorg.NewProperty("articleSection")

	// The number of words in the text of the Article.
	WordCount = schemaorg.NewProperty("wordCount")

	// The actual body of the article.
	ArticleBody = schemaorg.NewProperty("articleBody")

	// The page on which the work starts; for example "135" or "xiii".
	PageStart = schemaorg.NewProperty("pageStart")
)
