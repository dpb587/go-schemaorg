package liveblogposting

import "github.com/dpb587/go-schemaorg"

var (
	// The time when the live blog will begin covering the Event. Note that coverage
	// may begin before the Event's start time. The LiveBlogPosting may also be
	// created before coverage begins.
	CoverageStartTime = schemaorg.NewProperty("coverageStartTime")

	// The time when the live blog will stop covering the Event. Note that coverage
	// may continue after the Event concludes.
	CoverageEndTime = schemaorg.NewProperty("coverageEndTime")

	// An update to the LiveBlog.
	LiveBlogUpdate = schemaorg.NewProperty("liveBlogUpdate")
)
