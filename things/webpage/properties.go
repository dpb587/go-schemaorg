package webpage

import "github.com/dpb587/go-schemaorg"

var (
	// A set of links that can help a user understand and navigate a website
	// hierarchy.
	Breadcrumb = schemaorg.NewProperty("breadcrumb")

	// One of the more significant URLs on the page. Typically, these are the
	// non-navigation links that are clicked on the most.
	SignificantLink = schemaorg.NewProperty("significantLink")

	// A link related to this web page, for example to other related web pages.
	RelatedLink = schemaorg.NewProperty("relatedLink")

	// Date on which the content on this web page was last reviewed for accuracy
	// and/or completeness.
	LastReviewed = schemaorg.NewProperty("lastReviewed")

	// Indicates if this web page element is the main subject of the page.
	MainContentOfPage = schemaorg.NewProperty("mainContentOfPage")

	// People or organizations that have reviewed the content on this web page for
	// accuracy and/or completeness.
	ReviewedBy = schemaorg.NewProperty("reviewedBy")

	// Indicates the main image on the page.
	PrimaryImageOfPage = schemaorg.NewProperty("primaryImageOfPage")

	// The most significant URLs on the page. Typically, these are the
	// non-navigation links that are clicked on the most.
	SignificantLinks = schemaorg.NewProperty("significantLinks")

	// One of the domain specialities to which this web page's content applies.
	Specialty = schemaorg.NewProperty("specialty")
)
