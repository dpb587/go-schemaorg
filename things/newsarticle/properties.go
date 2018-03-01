package newsarticle

import "github.com/dpb587/go-schemaorg"

var (
	// The number of the column in which the NewsArticle appears in the print
	// edition.
	PrintColumn = schemaorg.NewProperty("printColumn")

	// The edition of the print product in which the NewsArticle appears.
	PrintEdition = schemaorg.NewProperty("printEdition")

	// If this NewsArticle appears in print, this field indicates the print section
	// in which the article appeared.
	PrintSection = schemaorg.NewProperty("printSection")

	// If this NewsArticle appears in print, this field indicates the name of the
	// page on which the article is found. Please note that this field is intended
	// for the exact page name (e.g. A5, B18).
	PrintPage = schemaorg.NewProperty("printPage")

	// A <a href="https://en.wikipedia.org/wiki/Dateline">dateline</a> is a brief
	// piece of text included in news articles that describes where and when the
	// story was written or filed though the date is often omitted. Sometimes only a
	// placename is provided.
	Dateline = schemaorg.NewProperty("dateline")
)
